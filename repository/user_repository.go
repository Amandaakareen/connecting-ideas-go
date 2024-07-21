package repository

import (
	"context"

	"github.com/example/entity"
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
