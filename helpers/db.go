package helpers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInsertOne(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
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

func DBInsertMany(collection *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
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

func DBFindOne(collection *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, filter, opts...)

	if result.Err() != nil {
		return nil, nil
	}
	var data map[string]interface{}
	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DBFindMany(collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.Find(ctx, filter, opts...)

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

func DBFindManyAndCount(collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		fmt.Println("failed to count records:", err)
		return nil, err
	}

	data, err := DBFindMany(collection, filter, opts...)
	result := map[string]interface{}{"count": count, "data": data}

	return result, err
}
