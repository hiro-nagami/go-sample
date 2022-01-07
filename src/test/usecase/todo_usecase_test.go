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
