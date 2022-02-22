package tests

import (
	"context"

	"github.com/chuckpreslar/emission"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	db.Init()
}

func InitApp() *models.App {
	return &models.App{
		Envs:         config.GetEnvs(),
		EventEmitter: emission.NewEmitter(),
	}
}

func DeleteTransactions() (*mongo.DeleteResult, error) {
	return db.GetCollection("transactions").DeleteMany(context.Background(), bson.M{"_id": bson.M{"$ne": nil}})
}

func CreateTransaction() string {
	txHash := "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af"
	transaction := models.Transaction{Hash: txHash}
	helpers.DB.InsertOne(db.GetCollection("transactions"), transaction)
	return txHash
}

func DeleteWatchList() (*mongo.DeleteResult, error) {
	return db.GetCollection("watch_list").DeleteMany(context.Background(), bson.M{"_id": bson.M{"$ne": nil}})
}
