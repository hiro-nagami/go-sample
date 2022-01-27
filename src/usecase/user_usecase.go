package usecase

import (
	"app/ent"
	"app/repository"
	"fmt"
)

type UserUseCase struct {
	Repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		Repo: repo,
	}
}

func (usecase *UserUseCase) CreateUser(name string, sex int) (*ent.User, error) {

	if name == "" {
		return nil, fmt.Errorf("%s", "Name is empty")
	}

	if sex < 0 || 4 < sex {
		return nil, fmt.Errorf("%s", "Sex is invalid")
	}

	user, err := usecase.Repo.CreateUser(name, sex)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}

func (usecase *UserUseCase) QueryUsers(id int) ([]*ent.User, error) {
	users, err := usecase.Repo.QueryUsers(id)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return users, nil
}
