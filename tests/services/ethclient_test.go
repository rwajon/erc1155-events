package tests

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/services"
	"github.com/rwajon/erc1155-events/tests"
	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func getClient() (*rpc.Client, *ethclient.Client) {
	envs := config.GetEnvs()
	rpcClient, _ := rpc.Dial(envs.RPCWebSocketURL)
	ethClient := ethclient.NewClient(rpcClient)
	return rpcClient, ethClient
}

func TestSubscribe(t *testing.T) {
	_, ethClient := getClient()
	logsCh := make(chan types.Log)
	tests.DeleteWatchList()
	assert.Error(t, services.Subscribe(*tests.InitApp(), nil, logsCh))
	assert.Error(t, services.Subscribe(*tests.InitApp(), ethClient, nil))
	assert.Error(t, services.Subscribe(*tests.InitApp(), ethClient, logsCh))
}

func TestGetBlock(t *testing.T) {
	rpcClient, _ := getClient()
	assert.NotNil(t, services.GetBlock(rpcClient, "latest"))
	assert.IsType(t, services.GetBlock(rpcClient, "latest"), &services.Block{})
}

func TestSaveTransaction(t *testing.T) {
	tests.DeleteTransactions()
	tx := models.Transaction{Hash: "test-tx-hash"}
	isCreated, err := services.SaveTransaction(tx)
	assert.True(t, isCreated)
	assert.Nil(t, err)
}

func TestGetBalance(t *testing.T) {
	rpcClient, ethClient := getClient()
	block := services.GetBlock(rpcClient, "latest")
	var transaction map[string]string
	json.Unmarshal(utils.Jsonify(block.Transactions[0]), &transaction)
	res := services.GetBalance(ethClient, transaction["from"], block.Number)
	assert.NotEqual(t, res, "")
}

func TestHandleNewLogDuplicatedError(t *testing.T) {
	rpcClient, _ := getClient()
	txHash := tests.CreateTransaction()
	log := types.Log{
		Address: common.HexToAddress(txHash),
		TxHash:  common.HexToHash("0x7f268357a8c2552623316e2562d90e642bb538e6"),
	}
	assert.Error(t, services.HandleNewLog(rpcClient, log), &services.Block{})
}
func TestHandleNewLogError(t *testing.T) {
	rpcClient, _ := getClient()
	log := types.Log{
		Address: common.HexToAddress("0x7f268357a8c2552623316e2562d90e642bb538e6"),
		TxHash:  common.HexToHash("0x7f268357a8c2552623316e2562d90e642bb538e6"),
	}
	assert.Error(t, services.HandleNewLog(rpcClient, log), &services.Block{})
}
