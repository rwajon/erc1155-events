package tests

import (
	"context"
	"testing"

	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = config.GetCollection("transactions")

func TestSave(t *testing.T) {
	transactionCollection.DeleteMany(context.Background(), bson.M{"_id": bson.M{"$ne": nil}})

	transaction := models.Transaction{
		Hash: "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af",
	}
	result, err := helpers.DBInsertOne(transactionCollection, transaction)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, *result, mongo.InsertOneResult{})

	result, _ = helpers.DBInsertOne(transactionCollection, nil)
	assert.Nil(t, result)

	var emptyData interface{}
	result, _ = helpers.DBInsertOne(transactionCollection, emptyData)
	assert.Nil(t, result)
}

func TestBulkSave(t *testing.T) {
	transactionCollection.DeleteMany(context.Background(), bson.M{"_id": bson.M{"$ne": nil}})

	transactions := []interface{}{
		models.Transaction{
			Hash: "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af",
		},
	}
	result, err := helpers.DBInsertMany(transactionCollection, transactions)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, *result, mongo.InsertManyResult{})

	result, _ = helpers.DBInsertMany(transactionCollection, nil)
	assert.Nil(t, result)

	var emptyData []interface{}
	result, _ = helpers.DBInsertMany(transactionCollection, emptyData)
	assert.Nil(t, result)
}
