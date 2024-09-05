package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongodrive "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Collection
}

func NewUserRepository(m *mongo.Collection) *UserRepository {
	return &UserRepository{client: m}
}

func (ur *UserRepository) Save(user entity.User) error {

	a, err := ur.client.InsertOne(context.Background(), user)
	fmt.Print(err)
	fmt.Print(a)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) ExistEmail(email string) error {
	var result entity.User
	fmt.Print(ur.client)

	err := ur.client.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&result)

	if err != nil && !errors.Is(err, mongodrive.ErrNoDocuments) {
		slog.Info("Teste", "flag", errors.Is(err, mongodrive.ErrNoDocuments))
		return err
	}

	if result.Email != "" {
		return errors.New("esse email já existe")
	}

	return nil

}

func (ur *UserRepository) FindByEmail(email string) (entity.User, error) {
	var result entity.User

	err := ur.client.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		return result, err
	}

	if result.Email == "" {
		return result, errors.New("esse usuário não existe")
	}

	return result, nil

}
