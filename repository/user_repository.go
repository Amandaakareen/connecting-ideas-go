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
	var result entity.User

	err := ur.client.FindOne(context.Background(), "users", bson.D{{Key: "email", Value: user.Email}}, &result)
	if err != nil {
		return err
	}

	if result.Email != "" {
		return errors.New("esse email já existe")
	}

	return nil

}
