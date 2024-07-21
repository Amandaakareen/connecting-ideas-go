package service

import (
	"github.com/example/entity"
	"github.com/example/repository"
	"gofr.dev/pkg/gofr"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{repository: ur}
}

func (us *UserService) Save(user entity.User, c *gofr.Context) error {
	err := us.repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
