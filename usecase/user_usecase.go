package usecase

import (
	"github.com/example/entity"
	"github.com/example/repository"
	"github.com/example/util"
	"gofr.dev/pkg/gofr"
)

type UserUseCase struct {
	repository *repository.UserRepository
}

func NewUserUsecase(ur *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repository: ur}
}

func (us *UserUseCase) Save(user entity.User, c *gofr.Context) error {
	err := us.repository.ExistEmail(user)
	if err != nil {
		return err
	}

	err = us.repository.ExistName(user)
	if err != nil {
		return err
	}

	util.EncryptPassword(&user)

	err = us.repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
