package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"app/utils/database"
	"context"
	"log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	var todo_id string
	db, _ := database.SetupDatabase()

	// log.Printf("Title: %s", input.Title)

	err := db.QueryRow("INSERT INTO todos(title, user_id, done) VALUES($1,$2,$3) RETURNING id", input.Title, input.UserID, false).Scan(&todo_id)
	if err != nil {
		log.Printf("Error - %s", err)
	}

	log.Printf("Todo ID: %s", todo_id)
	return todo_id, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	database.SetupDatabase()

	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
