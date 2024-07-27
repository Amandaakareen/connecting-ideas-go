package repository

import (
	"context"
	"errors"

	"github.com/example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr/datasource/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(m *mongo.Client) *UserRepository {
	return &UserRepository{client: m}
}

func (ur *UserRepository) Save(user entity.User) error {

	_, err := ur.client.InsertOne(context.Background(), "users", user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) ExistEmail(user entity.User) error {
	result := make([]entity.User, 0)

	err := ur.client.Find(context.Background(), "users", bson.D{{Key: "email", Value: user.Email}}, &result)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return errors.New("esse email já existe")
	}

	return nil

}

func (ur *UserRepository) ExistName(user entity.User) error {
	result := make([]entity.User, 0)

	err := ur.client.Find(context.Background(), "users", bson.D{{Key: "name", Value: user.Name}}, &result)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return errors.New("esse name já existe")
	}

	return nil

}
