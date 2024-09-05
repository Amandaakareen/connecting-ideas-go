package main

import (
	"log"

	"github.com/example/infra"
	"github.com/example/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	infra.ConnectMongo()
	infra.ConnectMinio()
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
