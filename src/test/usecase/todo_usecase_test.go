package usecase

import (
	"app/ent"
	"app/repository"
	"app/usecase"
	"testing"
)

type dummyTodoRepository struct {
	todos []*ent.Todo
}

func NewTodoRepository() repository.TodoRepository {
	return &dummyTodoRepository{}
}

func (repo *dummyTodoRepository) CreateTodo(title string, done bool, userId int) (*ent.Todo, error) {
	todo := &ent.Todo{
		ID:     1,
		Title:  title,
		Done:   done,
		UserID: userId,
	}
	repo.todos = append(repo.todos, todo)

	return todo, nil
}

func (repo *dummyTodoRepository) QueryTodos(id int) ([]*ent.Todo, error) {
	return repo.todos, nil
}

func TestTodoUseCase(t *testing.T) {

	t.Run("Add Todo", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(repository.NewTodoRepository())
		uc.CreateTodo("test1", false, 1)
		uc.CreateTodo("test2", false, 1)

		todos, err := uc.QueryTodos(1)

		if err != nil {
			t.Fatal("Couldn't create todo", err)
		}

		if len(todos) != 2 {
			t.Fatal("Couldn't create todo")
		}

		todo := todos[0]

		if todo.Title != "test" {
			t.Fatal("`Title` is wrong")
		}

		if todo.Done != false {
			t.Fatal("`Done` is wrong")
		}
	})

	t.Run("Add Todo", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(repository.NewTodoRepository())

		todo, err := uc.CreateTodo("", false, 1)
		uc.CreateTodo("test2", false, -1)

		todos, err := uc.QueryTodos(1)

		if err != nil {
			t.Fatal("Couldn't create todo", err)
		}

		if len(todos) != 2 {
			t.Fatal("Couldn't create todo")
		}

		todos, err := uc.QueryTodos(1)

		if len(todos) > 0 {
			t.Fatal("`Title` is wrong")
		}

		if todo.Done != false {
			t.Fatal("`Done` is wrong")
		}
	})
}
