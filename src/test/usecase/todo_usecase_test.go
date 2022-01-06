package usecase

import (
	"app/repository"
	"app/usecase"
	"testing"
)

func TestTodoUseCase(t *testing.T) {

	t.Run("Add Todo", func(t *testing.T) {
		usecase := usecase.NewTodoUseCase(repository.NewTodoRepository())
		usecase.CreateTodo("test", false, 1)

		todos, err := usecase.QueryTodos(1)

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
