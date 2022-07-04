package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoStart(uri string) *mongo.Collection {
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("example").Collection("users")
	println("Succesfully connected to mongodb")

	return collection
}
