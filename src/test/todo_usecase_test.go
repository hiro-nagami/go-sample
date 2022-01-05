package test

import (
	"app/usecase"
	"app/utils/database"
	"context"
	"testing"
)

func TestTodoUseCase(t *testing.T) {

	t.Run("Add Todo", func(t *testing.T) {
		client, err := database.GetEntClient()

		if err != nil {
			t.Fatal("Failed to create todo")
		}

		ctx := context.Background()
		usecase.CreateTodo("test", false, 1, ctx, client)
		defer client.Close()

		todos, err := usecase.QueryTodos(1, ctx, client)

		if err != nil {
			t.Fatal("Couldn't create todo", err)
		}

		todo := todos[0]

		if todo.Title != "test" {
			t.Fatal("Couldn't create todo")
		}

		if todo.Done != false {
			t.Fatal("Couldn't create todo")
		}

	})
}
