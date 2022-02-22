package db

import (
	"github.com/rwajon/erc1155-events/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var envs = config.GetEnvs()
var DBClient *mongo.Client

func Init() *mongo.Client {
	DBClient = config.ConnectDB()
	if DBClient != nil {
		Transaction.init()
		WatchList.init()
	}
	return DBClient
}

func GetCollection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection
	conn, _ := connstring.Parse(envs.DatabaseURL)
	collection = DBClient.Database(conn.Database).Collection(collectionName)

	return collection
}
