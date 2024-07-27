package controller

import (
	"github.com/example/entity"
	"github.com/example/usecase"
	"gofr.dev/pkg/gofr"
)

type UserController struct {
	service *usecase.UserUseCase
}

func NewUserController(us *usecase.UserUseCase) *UserController {
	return &UserController{service: us}
}

func (uc *UserController) Create(c *gofr.Context) (interface{}, error) {

	var user entity.User
	err := c.Request.Bind(&user)
	if err != nil {
		return nil, err

	}

	err = uc.service.Save(user, c)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
