package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInit() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect mongodb")
		return nil
	}

	fmt.Println("Connected to mongodb")
	return client
}

var Client *mongo.Client = DBInit()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var userCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return userCollection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var prodCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return prodCollection
}
