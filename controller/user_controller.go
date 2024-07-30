package controller

import (
	"log/slog"

	"github.com/example/dto"
	"github.com/example/usecase"
	"gofr.dev/pkg/gofr"
)

type UserController struct {
	usecase *usecase.UserUseCase
}

func NewUserController(us *usecase.UserUseCase) *UserController {
	return &UserController{usecase: us}
}

func (uc *UserController) Create(c *gofr.Context) (interface{}, error) {

	var user dto.UserCreate
	err := c.Request.Bind(&user)
	if err != nil {
		return nil, err

	}

	result, err := uc.usecase.Save(user)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return result, nil
}

func (uc *UserController) Login(c *gofr.Context) (interface{}, error) {

	var user dto.UserLogin
	err := c.Request.Bind(&user)
	if err != nil {
		return nil, err

	}

	token, err := uc.usecase.Login(user)

	if err != nil {
		return nil, err
	}

	return token, nil
}
