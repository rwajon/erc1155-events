package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/helpers"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection

func init() {
	db.Init()
	transactionCollection = db.GetCollection("transactions")
}

func TestSave(t *testing.T) {
	tests.DeleteTransactions()

	transaction := models.Transaction{
		Hash: "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af",
	}
	result, err := helpers.DB.InsertOne(transactionCollection, transaction)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, *result, mongo.InsertOneResult{})

	result, _ = helpers.DB.InsertOne(transactionCollection, nil)
	assert.Nil(t, result)

	var emptyData interface{}
	result, _ = helpers.DB.InsertOne(transactionCollection, emptyData)
	assert.Nil(t, result)
}

func TestBulkSave(t *testing.T) {
	tests.DeleteTransactions()

	transactions := []interface{}{
		models.Transaction{
			Hash: "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af",
		},
	}
	result, err := helpers.DB.InsertMany(transactionCollection, transactions)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, *result, mongo.InsertManyResult{})

	result, _ = helpers.DB.InsertMany(transactionCollection, nil)
	assert.Nil(t, result)

	var emptyData []interface{}
	result, _ = helpers.DB.InsertMany(transactionCollection, emptyData)
	assert.Nil(t, result)
}

func TestGetMany(t *testing.T) {
	tests.DeleteTransactions()

	transactions := []interface{}{
		models.Transaction{
			Hash: "0xdf18df8fe0150858d5bbbd149098fbd497adedefdfa91478960e71f07d0019af",
		},
	}

	helpers.DB.InsertMany(transactionCollection, transactions)

	result, err := helpers.DB.FindMany(transactionCollection, nil)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result), 0)

	tests.DeleteTransactions()
	result, err = helpers.DB.FindMany(transactionCollection, nil)
	assert.Equal(t, len(result), 0)
	assert.Nil(t, err)
}

func TestDeleteOne(t *testing.T) {
	tests.DeleteTransactions()

	transactions := []interface{}{
		models.Transaction{
			Hash: "tx-hash",
		},
	}

	helpers.DB.InsertMany(transactionCollection, transactions)

	result, err := helpers.DB.DeleteOne(transactionCollection, bson.M{"hash": "tx-hash"}, nil)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, result.DeletedCount, int64(0))

	result, err = helpers.DB.DeleteOne(transactionCollection, bson.M{"hash": "tx-hash"}, nil)
	assert.Equal(t, result.DeletedCount, int64(0))
	assert.Nil(t, err)
}

func TestUpdateOne(t *testing.T) {
	tests.DeleteTransactions()

	transactions := []interface{}{
		models.Transaction{
			Hash: "tx-hash",
		},
	}

	helpers.DB.InsertMany(transactionCollection, transactions)

	result, err := helpers.DB.UpdateOne(transactionCollection, bson.M{"hash": "tx-hash"},
		bson.D{{"$set", bson.D{{"hash", "tx-new-hash"}}}}, nil)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, result.ModifiedCount, int64(0))

	result, err = helpers.DB.UpdateOne(transactionCollection, bson.M{"hash": "tx-hash"},
		bson.D{{"$set", bson.D{{"hash", "tx-new-hash"}}}}, nil)
	assert.Equal(t, result.ModifiedCount, int64(0))
	assert.Nil(t, err)
}
