package infra

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectMongo() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://admin:adimin@localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	db := client.Database("connecting")

	collection := db.Collection("users")

	Collection = collection

}
