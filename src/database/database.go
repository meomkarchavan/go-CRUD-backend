package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection

const db = "training_db"

var userCollection = "users"
var postCollection = "post"

var url = "mongodb://localhost:27017/"

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func createConnection(url string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	clientOptions := options.Client().ApplyURI(url)

	ctx, cancel := context.WithTimeout(context.Background(),
		60*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel, err
}
