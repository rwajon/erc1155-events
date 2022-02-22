package services

import (
	"context"
	"errors"
	"strings"

	"fmt"
	"log"
	"math"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/constants"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type Block struct {
	Hash         string `json:"hash"`
	Number       string `json:"number"`
	Timestamp    string `json:"timestamp"`
	Transactions []Transaction
}

type Transaction struct {
	From     string `json:"from"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Hash     string `json:"hash"`
	To       string `json:"to"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

var envs config.Env = config.GetEnvs()

func ListenToERC1155Events(app models.App) error {
	rpcClient, err := rpc.Dial(envs.RPCWebSocketURL)
	ethClient := ethclient.NewClient(rpcClient)

	if err != nil {
		log.Println("RPC client Error:", err.Error())
		return err
	}

	log.Println("RPC client connected", rpcClient)

	logsCh := make(chan types.Log)

	go Subscribe(app, ethClient, logsCh)

	for log := range logsCh {
		go HandleNewLog(rpcClient, log)
	}

	return nil
}

func GetAddressesToWatch() []common.Address {
	var contractAddresses []common.Address
	if data, err := db.WatchList.GetMany(nil); err == nil {
		for _, v := range data {
			address := fmt.Sprintf("%v", v["address"])
			contractAddresses = append(contractAddresses, common.HexToAddress(string(address)))
		}
	}
	return contractAddresses
}

func HandleNewLog(client *rpc.Client, log types.Log) error {
	fmt.Println("new log: ", log.TxHash.String())
	result, _ := db.Transaction.GetOne(bson.M{
		"hash": bson.M{"$regex": log.TxHash.String(), "$options": "im"},
	})
	if result != nil {
		return errors.New("duplicated transaction: " + string(utils.Jsonify(result)))
	}
	tx := GetTransaction(client, log)
	if tx == nil {
		return errors.New("failed to get transaction")
	}
	_, err := SaveTransaction(*tx)
	return err
}

func GetTransaction(client *rpc.Client, log types.Log) *models.Transaction {
	block := GetBlock(client, log.BlockHash.Hex())

	if block == nil {
		return nil
	}

	for _, tx := range block.Transactions {
		if strings.EqualFold(tx.Hash, log.TxHash.String()) {
			ethClient := ethclient.NewClient(client)
			transaction := &models.Transaction{
				BlockNumber:     block.Number,
				BlockHash:       block.Hash,
				Timestamp:       utils.HexToInt(block.Timestamp),
				Date:            time.Unix(utils.HexToInt(block.Timestamp), 0),
				From:            tx.From,
				Gas:             utils.HexToFloat(tx.Gas) / math.Pow10(18),
				GasPrice:        utils.HexToFloat(tx.GasPrice) / math.Pow10(18),
				Hash:            tx.Hash,
				To:              tx.To,
				Type:            uint8(utils.HexToInt(tx.Type)),
				Value:           utils.HexToFloat(tx.Value) / math.Pow10(18),
				SenderBalance:   GetBalance(ethClient, tx.From, block.Number),
				ReceiverBalance: GetBalance(ethClient, tx.To, block.Number),
				ContractAddress: log.Address.String(),
			}
			return transaction
		}
	}
	return nil
}

func GetBlock(client *rpc.Client, blockHash string) *Block {
	var block Block

	if client == nil {
		fmt.Println("client should not be nil")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := client.CallContext(ctx, &block, func() string {
		if blockHash == "latest" {
			return "eth_getBlockByNumber"
		}
		return "eth_getBlockByHash"
	}(), blockHash, true)

	if err != nil {
		fmt.Println("can't get block:", blockHash, err)
		return nil
	}

	return &block
}

func SaveTransaction(tx models.Transaction) (bool, error) {
	_, err := db.Transaction.Save(tx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetBalance(client *ethclient.Client, address string, blockNumber string) float64 {
	if client == nil {
		fmt.Println("client should not be nil")
		return 0
	}
	_address := common.HexToAddress(address)
	_blockNumber := common.HexToAddress(blockNumber).Hash().Big()
	balance, err := client.BalanceAt(context.Background(), _address, _blockNumber)

	if err != nil {
		fmt.Println("can't get balance from address:", address)
		return 0
	}

	return utils.StringToFloat(balance.String()) / math.Pow10(18)
}

func Subscribe(app models.App, client *ethclient.Client, logsCh chan types.Log) error {
	if client == nil || logsCh == nil {
		return errors.New("client and log channel should not be nil")
	}

	contractAddresses := GetAddressesToWatch()

	if len(contractAddresses) == 0 {
		return errors.New("no contract address to watch")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := ethereum.FilterQuery{Addresses: contractAddresses}
	ethSub, err := client.SubscribeFilterLogs(ctx, query, logsCh)

	if err != nil {
		return errors.New("subscribe error: " + err.Error())
	}

	app.EventEmitter.On(constants.OnWatchListChange, func(data interface{}) {
		ethSub.Unsubscribe()
		go Subscribe(app, client, logsCh)
	})

	fmt.Println("connection lost: ", <-ethSub.Err())
	return <-ethSub.Err()
}
