package usecase

import (
	"app/ent"
	"app/ent/todo"
	"app/utils/database"
	"context"
	"fmt"
	"log"
)

func CreateTodo(title string, done bool, userId string, ctx context.Context) (*ent.Todo, error) {
	client, err := database.GetEntClient()

	todo, err := client.Todo.
		Create().
		SetTitle(title).
		SetDone(done).
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", todo)
	return todo, nil
}

func QueryUser(id int, ctx context.Context) ([]*ent.Todo, error) {
	client, err := database.GetEntClient()

	todos, err := client.Todo.
		Query().
		Where(todo.IDEQ(id)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", todos)
	return todos, nil
}
