package services

import (
	"context"

	"fmt"
	"log"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/utils"
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

func ListenToERC1155Events() error {
	rpcClient, err := rpc.Dial(envs.RPCWebSocketURL)
	ethClient := ethclient.NewClient(rpcClient)

	if err != nil {
		log.Println("RPC Error:", err.Error())
		return err
	}

	log.Println("RPC connected", rpcClient)
	log.Println("ethClient connected", ethClient)

	blocksCh := make(chan Block)

	go SubscribeBlocks(rpcClient, blocksCh)

	for block := range blocksCh {
		if len(block.Transactions) > 0 {
			go SaveBlockTransactions(ethClient, block)
		} else {
			go SaveBlockTransactions(ethClient, GetBlock(rpcClient, block.Number))
		}
	}
	return nil
}

func GetBlock(client *rpc.Client, blockNumber string) Block {
	var block Block

	if client == nil {
		fmt.Println("client should not be nil")
		return block
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := client.CallContext(ctx, &block, "eth_getBlockByNumber", blockNumber, true); err != nil {
		fmt.Println("can't get block:", block.Number, err)
	}

	return block
}

func SaveBlockTransactions(client *ethclient.Client, block Block) []models.Transaction {
	if client == nil {
		fmt.Println("client should not be nil")
		return nil
	}

	log.Println("new block:", block.Number)
	var transactions []models.Transaction

	for _, tx := range block.Transactions {
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
			SenderBalance:   GetBalance(client, tx.From, block.Number),
			ReceiverBalance: GetBalance(client, tx.To, block.Number),
			ContractAddress: GetContract(client, tx.To, block.Number),
		}

		transactions = append(transactions, *transaction)
	}

	db.Transaction.BulkSave(transactions, &options.InsertManyOptions{
		Ordered: func() *bool { v := false; return &v }(),
	})

	return transactions
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

func GetContract(client *ethclient.Client, address string, blockNumber string) string {
	if client == nil {
		fmt.Println("client should not be nil")
		return ""
	}
	var contract string
	_address := common.HexToAddress(address)
	_blockNumber := common.HexToAddress(blockNumber).Hash().Big()
	bytecode, err := client.CodeAt(context.Background(), _address, _blockNumber)

	if err != nil {
		fmt.Println("can't get contract code from address:", _address)
	} else if len(bytecode) > 0 {
		contract = _address.String()
	}
	return contract
}

func SubscribeBlocks(client *rpc.Client, blocksCh chan Block) error {
	if client == nil || blocksCh == nil {
		fmt.Println("client and subscribe channel should not be nil")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ethSub, err := client.EthSubscribe(ctx, blocksCh, "newHeads")
	if err != nil {
		fmt.Println("subscribe error:", err)
		return err
	}

	var block Block
	err = client.CallContext(ctx, &block, "eth_getBlockByNumber", "latest", true)

	if err != nil {
		fmt.Println("can't get latest block:", err)
		return err
	}

	// TODO: remove Unsubscribe
	// ethSub.Unsubscribe()
	blocksCh <- block

	fmt.Println("connection lost: ", <-ethSub.Err())
	return nil
}
