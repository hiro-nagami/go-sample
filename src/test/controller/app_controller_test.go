package controller

import (
	"app/ent"
	"app/repository"
	"app/usecase"
	"testing"
)

type dummyTodoRepository struct {
	todos []*ent.Todo
	count int
}

func NewDummyTodoRepository() repository.TodoRepository {
	return &dummyTodoRepository{
		todos: []*ent.Todo{},
		count: 0,
	}
}

func (repo *dummyTodoRepository) CreateTodo(title string, done bool, userId int) (*ent.Todo, error) {
	repo.count++

	todo := &ent.Todo{
		ID:     repo.count,
		Title:  title,
		Done:   done,
		UserID: userId,
	}
	repo.todos = append(repo.todos, todo)

	return todo, nil
}

func (repo *dummyTodoRepository) QueryTodos(id int) ([]*ent.Todo, error) {
	todos := []*ent.Todo{}

	for i := range repo.todos {
		if repo.todos[i].ID == id {
			todos = append(todos, repo.todos[i])
		}
	}

	return todos, nil
}

func (repo *dummyTodoRepository) QueryTodosByUserID(userId int) ([]*ent.Todo, error) {
	todos := []*ent.Todo{}

	for i := range repo.todos {
		if repo.todos[i].UserID == userId {
			todos = append(todos, repo.todos[i])
		}
	}

	return todos, nil
}

func TestTodoUseCase(t *testing.T) {

	t.Run("Add Todo", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(NewDummyTodoRepository())
		uc.CreateTodo("test1", false, 1)
		uc.CreateTodo("test2", true, 1)

		todos, err := uc.QueryTodos(2)

		if err != nil {
			t.Fatal("Couldn't create todo", err)
		}

		if len(todos) != 1 {
			t.Fatal("Didn't match todos len")
		}

		todo := todos[0]

		if todo.Title != "test2" {
			t.Fatal("`Title` is wrong")
		}

		if todo.Done != true {
			t.Fatal("`Done` is wrong")
		}
	})

	t.Run("Add Todo", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(NewDummyTodoRepository())
		uc.CreateTodo("test1", false, 1)
		uc.CreateTodo("test2", false, 1)

		todos, err := uc.QueryTodosByUserID(1)

		if err != nil {
			t.Fatal("Couldn't create todo", err)
		}

		if len(todos) != 2 {
			t.Fatal("Didn't match todos len")
		}

		todo := todos[0]

		if todo.Title != "test1" {
			t.Fatal("`Title` is wrong")
		}

		if todo.Done != false {
			t.Fatal("`Done` is wrong")
		}
	})

	t.Run("Failed to create todo by empty tilte", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(NewDummyTodoRepository())

		todo, err := uc.CreateTodo("", false, 1)

		if todo != nil || err == nil {
			t.Fatal("UseCase couldn't classify invalid cases")
		}

		todos, err := uc.QueryTodos(1)

		if len(todos) > 0 {
			t.Fatal("`Title` is wrong")
		}
	})

	t.Run("Failed to create todo by invalid userid", func(t *testing.T) {
		uc := usecase.NewTodoUseCase(NewDummyTodoRepository())

		todo, err := uc.CreateTodo("test", false, -1)

		if todo != nil || err == nil {
			t.Fatal("UseCase couldn't classify invalid cases")
		}

		todos, err := uc.QueryTodos(1)

		if len(todos) > 0 {
			t.Fatal("`Title` is wrong")
		}
	})
}
