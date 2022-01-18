package usecase

import (
	"app/ent"
	"app/repository"
	"fmt"
)

type TodoUseCase struct {
	Repo repository.TodoRepository
}

func NewTodoUseCase(repo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		Repo: repo,
	}
}

func (usecase *TodoUseCase) CreateTodo(title string, done bool, userId int) (*ent.Todo, error) {

	if title == "" {
		return nil, fmt.Errorf("%s", "Title is empty")
	}

	if userId <= 0 {
		return nil, fmt.Errorf("%s", "User ID is wrong")
	}

	todo, err := usecase.Repo.CreateTodo(title, done, userId)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return todo, nil
}

func (usecase *TodoUseCase) QueryTodos(id int) ([]*ent.Todo, error) {
	todos, err := usecase.Repo.QueryTodos(id)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return todos, nil
}

func (usecase *TodoUseCase) QueryTodosByUserID(userId int) ([]*ent.Todo, error) {
	todos, err := usecase.Repo.QueryTodosByUserID(userId)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return todos, nil
}
