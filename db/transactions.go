package db

import (
	"context"
	"log"

	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITransaction interface {
	Save(data models.Transaction, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	BulkSave(data []models.Transaction, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
}

type transaction models.Transaction

var transactionCollection *mongo.Collection = config.GetCollection("transactions")

var Transaction = new(transaction)

func (tx *transaction) createIndexes() {
	indexModel, err := transactionCollection.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "hash", Value: 1}},
			Options: options.Index().SetUnique(true),
		})

	if err != nil {
		log.Fatalf("can't create transaction index: %+v", err)
	}
	log.Println("transaction index", indexModel)
}

func (tx *transaction) Save(data models.Transaction, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return helpers.DBInsertOne(transactionCollection, data, opts...)
}

func (tx *transaction) BulkSave(data []models.Transaction, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var transactions []interface{}
	for _, tx := range data {
		transactions = append(transactions, tx)
	}
	return helpers.DBInsertMany(transactionCollection, transactions, opts...)
}
