package main

import (
	"github.com/example/controller"
	"github.com/example/infra"
	"github.com/example/repository"
	"github.com/example/service"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	db := infra.Connect()
	app.AddMongo(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app.POST("/", userController.Create)
	app.Run()
}
