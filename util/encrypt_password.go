package util

import (
	"github.com/example/entity"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(user *entity.User) error {

	result, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(result)

	return nil

}
