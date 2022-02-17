package tests

import (
	"context"

	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection = config.GetCollection("transactions")

func DeleteTransactions() (*mongo.DeleteResult, error) {
	return transactionCollection.DeleteMany(context.Background(), bson.M{"_id": bson.M{"$ne": nil}})
}

func CreateTransaction() string {
	txHash := "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af"
	transaction := models.Transaction{Hash: txHash}
	helpers.DBInsertOne(transactionCollection, transaction)
	return txHash
}
