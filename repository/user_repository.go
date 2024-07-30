package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/example/entity"
	"go.mongodb.org/mongo-driver/bson"
	mongodrive "go.mongodb.org/mongo-driver/mongo"
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

func (ur *UserRepository) ExistEmail(email string) error {
	var result entity.User

	err := ur.client.FindOne(context.Background(), "users", bson.D{{Key: "email", Value: email}}, &result)
	if err != nil && !errors.Is(err, mongodrive.ErrNoDocuments) {
		slog.Info("Teste", "flag", errors.Is(err, mongodrive.ErrNoDocuments))
		return err
	}

	if result.Email != "" {
		return errors.New("esse email já existe")
	}

	return nil

}

func (ur *UserRepository) ExistName(name string) error {
	var result entity.User

	err := ur.client.FindOne(context.Background(), "users", bson.D{{Key: "name", Value: name}}, &result)
	if err != nil && !errors.Is(err, mongodrive.ErrNoDocuments) {
		return err
	}

	if result.Name != "" {
		return errors.New("esse name já existe")
	}

	return nil

}

func (ur *UserRepository) FindByEmail(email string) (entity.User, error) {
	var result entity.User

	err := ur.client.FindOne(context.Background(), "users", bson.D{{Key: "email", Value: email}}, &result)
	if err != nil {
		return result, err
	}

	if result.Email == "" {
		return result, errors.New("esse usuário não existe")
	}

	return result, nil

}
