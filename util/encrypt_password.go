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

func CheckPassword(encryptPassword, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(loginPassword))
	if err != nil {
		return err
	}
	return nil
}
