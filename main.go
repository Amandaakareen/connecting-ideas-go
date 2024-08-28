package main

import (
	"log/slog"
	"net/http"

	"github.com/example/controller"
	"github.com/example/infra"
	"github.com/example/repository"
	"github.com/example/usecase"
	"github.com/gorilla/mux"
)

func main() {
	forever := make(chan struct{})

	infra.ConnectMongo()
	infra.ConnectMinio()

	userRepository := repository.NewUserRepository(infra.Mongo)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	mux := mux.NewRouter()

	mux.HandleFunc("/", userController.Create).Methods(http.MethodPost)
	mux.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	mux.HandleFunc(" /foto", userController.AddPhoto).Methods(http.MethodPost)

	go http.ListenAndServe(":8000", mux)

	slog.Info("Server is ready on port 8000")

	<-forever
}
