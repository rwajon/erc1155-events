package db

import (
	"context"
	"log"

	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transaction models.Transaction

var transactionCollection *mongo.Collection
var Transaction = new(transaction)

func (tx *transaction) init() {
	transactionCollection = GetCollection("transactions")
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
	return helpers.DB.InsertOne(transactionCollection, data, opts...)
}

func (tx *transaction) BulkSave(data []models.Transaction, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var transactions []interface{}
	for _, tx := range data {
		transactions = append(transactions, tx)
	}
	return helpers.DB.InsertMany(transactionCollection, transactions, opts...)
}

func (tx *transaction) GetOne(filter interface{}, opts ...*options.FindOneOptions) (map[string]interface{}, error) {
	return helpers.DB.FindOne(transactionCollection, filter, opts...)
}

func (tx *transaction) GetMany(filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error) {
	return helpers.DB.FindMany(transactionCollection, filter, opts...)
}

func (tx *transaction) GetManyAndCount(filter interface{}, opts ...*options.FindOptions) (*helpers.DBFindManyAndCount, error) {
	return helpers.DB.FindManyAndCount(transactionCollection, filter, opts...)
}
