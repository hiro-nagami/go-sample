package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"app/usecase"
	"app/utils/database"
	"context"
	"log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	client, err := database.GetEntClient()

	if err != nil {
		return nil, err
	}

	todo, err := usecase.CreateTodo(input.Title, false, input.UserID, ctx, client)
	defer client.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rTodo := &model.Todo{
		ID:     todo.ID,
		Title:  todo.Title,
		Done:   todo.Done,
		UserID: todo.UserID,
	}

	log.Printf("Todo ID: %d", todo.ID)

	return rTodo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	client, err := database.GetEntClient()

	if err != nil {
		return nil, err
	}

	todos, _ := usecase.QueryTodos(1, ctx, client)
	defer client.Close()

	rTodos := []*model.Todo{}

	for _, todo := range todos {
		rTodo := &model.Todo{
			ID:     todo.ID,
			Title:  todo.Title,
			Done:   todo.Done,
			UserID: todo.UserID,
		}
		rTodos = append(rTodos, rTodo)
	}

	return rTodos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
