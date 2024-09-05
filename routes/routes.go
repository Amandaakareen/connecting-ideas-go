package routes

import (
	"github.com/example/controller"
	"github.com/example/infra"
	"github.com/example/repository"
	"github.com/example/usecase"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	userRepository := repository.NewUserRepository(infra.Collection)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	r.POST("/", userController.Create)
	r.POST("/foto/:id", userController.Login)
	r.POST("/login", userController.AddPhoto)
}
