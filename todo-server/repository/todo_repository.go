package repository

import (
	"app/ent"
	"app/ent/todo"
	"app/utils/database"
	"context"
	"fmt"
)

type TodoRepository interface {
	CreateTodo(title string, done bool, userId int) (*ent.Todo, error)
	QueryTodos(id int) ([]*ent.Todo, error)
	QueryTodosByUserID(userId int) ([]*ent.Todo, error)
}

type todoRepository struct {
	client  *ent.Client
	context context.Context
}

func NewTodoRepository() TodoRepository {
	client, err := database.GetEntClient()

	if err != nil {
		return nil
	}

	return &todoRepository{
		client:  client,
		context: context.Background(),
	}
}

func (repo *todoRepository) CreateTodo(title string, done bool, userId int) (*ent.Todo, error) {
	todo, err := repo.client.Todo.
		Create().
		SetTitle(title).
		SetDone(done).
		SetUserID(userId).
		Save(repo.context)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	defer repo.client.Close()
	return todo, nil
}

func (repo *todoRepository) QueryTodos(id int) ([]*ent.Todo, error) {
	client, err := database.GetEntClient()
	context := context.Background()

	todos, err := client.Todo.
		Query().
		Where(todo.IDEQ(id)).
		All(context)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	defer repo.client.Close()
	return todos, nil
}

func (repo *todoRepository) QueryTodosByUserID(userId int) ([]*ent.Todo, error) {
	client, err := database.GetEntClient()
	context := context.Background()

	todos, err := client.Todo.
		Query().
		Where(todo.UserID(userId)).
		All(context)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	defer repo.client.Close()
	return todos, nil
}
