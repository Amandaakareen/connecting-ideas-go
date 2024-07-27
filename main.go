package main

import (
	"github.com/example/controller"
	"github.com/example/infra"
	"github.com/example/repository"
	"github.com/example/usecase"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	db := infra.Connect()
	app.AddMongo(db)

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	app.POST("/", userController.Create)
	app.Run()
}
