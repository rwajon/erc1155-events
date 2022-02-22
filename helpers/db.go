package helpers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBFindManyAndCount struct {
	Data  []map[string]interface{}
	Count int64
}

type database struct{}

var DB = new(database)

func (db *database) InsertOne(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	if document == nil {
		fmt.Println("document to save is required")
		return nil, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, document, opts...)

	if err != nil {
		fmt.Println("failed to save:", err)
		return nil, err
	}
	return result, err
}

func (db *database) InsertMany(collection *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if documents == nil {
		fmt.Println("documents to save are required")
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertMany(ctx, documents, opts...)

	if err != nil {
		fmt.Println("failed to bulk save:", err)
		return nil, err
	}
	return result, err
}

func (db *database) FindOne(collection *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, func() interface{} {
		if filter == nil {
			return bson.M{}
		}
		return filter
	}(), opts...)

	if result.Err() != nil {
		return nil, nil
	}
	var data map[string]interface{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (db *database) FindMany(collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.Find(ctx, func() interface{} {
		if filter == nil {
			return bson.M{}
		}
		return filter
	}(), opts...)

	if err != nil {
		fmt.Println("failed to find:", err)
		return nil, err
	}

	defer result.Close(ctx)

	var data []map[string]interface{}

	if err = result.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (db *database) FindManyAndCount(collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) (*DBFindManyAndCount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		fmt.Println("failed to count records:", err)
		return nil, err
	}

	data, err := DB.FindMany(collection, filter, opts...)
	result := &DBFindManyAndCount{Count: count, Data: data}

	return result, err
}

func (db *database) UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update, opts...)

	if err != nil {
		fmt.Println("failed to update record:", err)
		return nil, err
	}

	return result, nil
}

func (db *database) DeleteOne(collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, func() interface{} {
		if filter == nil {
			return bson.M{}
		}
		return filter
	}(), opts...)

	if err != nil {
		fmt.Println("failed to delete record:", err)
		return nil, err
	}

	return result, nil
}
