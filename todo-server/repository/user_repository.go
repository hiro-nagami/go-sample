package repository

import (
	"app/ent"
	"app/ent/user"
	"app/utils/database"
	"context"
	"fmt"
)

type UserRepository interface {
	CreateUser(name string, sex int) (*ent.User, error)
	QueryUsers(id int) ([]*ent.User, error)
}

type userRepository struct {
	client  *ent.Client
	context context.Context
}

func NewUserRepository() UserRepository {
	client, err := database.GetEntClient()

	if err != nil {
		return nil
	}

	return &userRepository{
		client:  client,
		context: context.Background(),
	}
}

func (repo *userRepository) CreateUser(name string, sex int) (*ent.User, error) {
	user, err := repo.client.User.
		Create().
		SetName(name).
		SetSex(sex).
		Save(repo.context)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	defer repo.client.Close()
	return user, nil
}

func (repo *userRepository) QueryUsers(id int) ([]*ent.User, error) {
	client, err := database.GetEntClient()
	context := context.Background()

	users, err := client.User.
		Query().
		Where(user.IDEQ(id)).
		All(context)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	defer repo.client.Close()
	return users, nil
}
