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

type IWatchList interface {
	Save(data models.WatchList, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	BulkSave(data []models.WatchList, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	GetOne(filter interface{}, opts ...*options.FindOneOptions) (map[string]interface{}, error)
	GetMany(filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error)
	GetManyAndCount(filter interface{}, opts ...*options.FindOptions) (map[string]interface{}, error)
	UpdateOne(filter interface{}, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(filter interface{}, data interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type watchList models.WatchList

var watchListCollection *mongo.Collection = config.GetCollection("watch_list")
var WatchList = new(watchList)

func (tx *watchList) createIndexes() {
	indexModel, err := watchListCollection.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "address", Value: 1}},
			Options: options.Index().SetUnique(true),
		})

	if err != nil {
		log.Fatalf("can't create watch list index: %+v", err)
	}
	log.Println("watch list index", indexModel)
}

func (tx *watchList) Save(data models.WatchList, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return helpers.DBInsertOne(watchListCollection, data, opts...)
}

func (tx *watchList) BulkSave(data []models.WatchList, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var watchList []interface{}
	for _, tx := range data {
		watchList = append(watchList, tx)
	}
	return helpers.DBInsertMany(watchListCollection, watchList, opts...)
}

func (tx *watchList) GetOne(filter interface{}, opts ...*options.FindOneOptions) (map[string]interface{}, error) {
	return helpers.DBFindOne(watchListCollection, filter, opts...)
}

func (tx *watchList) GetMany(filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error) {
	return helpers.DBFindMany(watchListCollection, filter, opts...)
}

func (tx *watchList) GetManyAndCount(filter interface{}, opts ...*options.FindOptions) (*models.DBFindManyAndCount, error) {
	return helpers.DBFindManyAndCount(watchListCollection, filter, opts...)
}

func (tx *watchList) UpdateOne(filter interface{}, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return helpers.DBUpdateOne(watchListCollection, filter, data, opts...)
}

func (tx *watchList) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return helpers.DBDeleteOne(watchListCollection, filter, opts...)
}
