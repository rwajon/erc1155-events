package services

import (
	"context"
	"encoding/json"

	"fmt"
	"log"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	Number       string `json:"number"`
	Timestamp    string `json:"timestamp"`
	Transactions []interface{}
}

func ListenToERC1155Events() error {
	envs := config.GetEnvs()
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
			SaveBlockTransactions(ethClient, block)
		} else {
			go func() {
				_block := GetBlock(rpcClient, block.Number)
				SaveBlockTransactions(ethClient, _block)
			}()
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

	var transactions, _transactions []models.Transaction

	if err := json.Unmarshal(utils.InterfaceToJson(block.Transactions), &_transactions); err == nil {
		for i, tx := range _transactions {
			_tx := &models.Transaction{
				BlockNumber: block.Number,
				BlockHash:   block.Hash,
				Timestamp:   fmt.Sprintf("%d", utils.HexToInt(block.Timestamp)),
				Date:        time.Unix(utils.HexToInt(block.Timestamp), 0).String(),
				From:        tx.From,
				Gas:         fmt.Sprintf("%.18f", utils.HexToFloat(tx.Gas)/math.Pow10(18)),
				GasPrice:    fmt.Sprintf("%.18f", utils.HexToFloat(tx.GasPrice)/math.Pow10(18)),
				Hash:        tx.Hash,
				To:          tx.To,
				Type:        fmt.Sprintf("%d", utils.HexToInt(tx.Type)),
				Value:       fmt.Sprintf("%.18f", utils.HexToFloat(tx.Value)/math.Pow10(18)),
			}

			if i < 1 { //TODO: remove condition
				_tx.SenderBalance = GetBalance(client, tx.From, block.Number)
				_tx.ReceiverBalance = GetBalance(client, tx.To, block.Number)
				_tx.ContractAddress = GetContract(client, tx.To, block.Number)
			}
			transactions = append(transactions, *_tx)
		}

		//TODO: add save to database implementation
		if len(transactions) > 0 {
			utils.WriteToFile("transactions.log", string(utils.InterfaceToJson(transactions)))
		}
	} else {
		return nil
	}

	return transactions
}

func GetBalance(client *ethclient.Client, address string, blockNumber string) string {
	if client == nil {
		fmt.Println("client should not be nil")
		return ""
	}
	_address := common.HexToAddress(address)
	_blockNumber := common.HexToAddress(blockNumber).Hash().Big()
	balance, err := client.BalanceAt(context.Background(), _address, _blockNumber)

	if err != nil {
		fmt.Println("can't get balance from address:", address)
		return ""
	}

	return fmt.Sprintf("%.18f", utils.StringToFloat(balance.String())/math.Pow10(18))
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
