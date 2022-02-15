package tests

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/services"
	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func getClient() (*rpc.Client, *ethclient.Client) {
	envs := config.GetEnvs()
	rpcClient, _ := rpc.Dial(envs.RPCWebSocketURL)
	ethClient := ethclient.NewClient(rpcClient)
	return rpcClient, ethClient
}

func TestSubscribeBlocks(t *testing.T) {
	rpcClient, _ := getClient()
	assert.Nil(t, services.SubscribeBlocks(rpcClient, nil))
}

func TestGetBlock(t *testing.T) {
	rpcClient, _ := getClient()
	assert.NotNil(t, services.GetBlock(rpcClient, "latest"))
	assert.IsType(t, services.GetBlock(rpcClient, "latest"), services.Block{})
}

func TestSaveBlockTransactions(t *testing.T) {
	rpcClient, ethClient := getClient()
	block := services.GetBlock(rpcClient, "latest")
	assert.NotNil(t, services.SaveBlockTransactions(ethClient, block))
	assert.IsType(t, services.SaveBlockTransactions(ethClient, block), []models.Transaction{})
	assert.Nil(t, services.SaveBlockTransactions(nil, block))
}

func TestGetBalance(t *testing.T) {
	rpcClient, ethClient := getClient()
	block := services.GetBlock(rpcClient, "latest")
	var transaction map[string]string
	json.Unmarshal(utils.InterfaceToJson(block.Transactions[0]), &transaction)
	res := services.GetBalance(ethClient, transaction["from"], block.Number)
	assert.NotEqual(t, res, "")
}
