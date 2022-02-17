package helpers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSave(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
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

func DBBulkSave(collection *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
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
