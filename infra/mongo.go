package infra

import (
	"gofr.dev/pkg/gofr/datasource/mongo"
)

func Connect() *mongo.Client {

	db := mongo.New(mongo.Config{URI: "mongodb://admin:adimin@localhost:27017", Database: "connecting"})

	return db

}
