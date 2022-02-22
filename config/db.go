package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var envs Env = GetEnvs()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(envs.DatabaseURL))

	if err != nil {
		log.Println(err)
		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Connected to MongoDB")

	return client
}
