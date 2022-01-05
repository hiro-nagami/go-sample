package usecase

import (
	"app/ent"
	"app/ent/todo"
	"context"
	"fmt"
)

func CreateTodo(title string, done bool, userId int, ctx context.Context, client *ent.Client) (*ent.Todo, error) {
	todo, err := client.Todo.
		Create().
		SetTitle(title).
		SetDone(done).
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return todo, nil
}

func QueryTodos(id int, ctx context.Context, client *ent.Client) ([]*ent.Todo, error) {
	todos, err := client.Todo.
		Query().
		Where(todo.IDEQ(id)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return todos, nil
}
