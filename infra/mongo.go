package infra

import (
	"log"

	"gofr.dev/pkg/gofr/datasource/mongo"
)

var Mongo *mongo.Client

func ConnectMongo() {
	Mongo = mongo.New(mongo.Config{URI: "mongodb://admin:adimin@localhost:27017", Database: "connecting"})

	if Mongo == nil {
		log.Fatalf("Erro ao conectar com o MongoDB ")
	}
}
