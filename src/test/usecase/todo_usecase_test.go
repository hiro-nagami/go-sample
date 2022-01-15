package usecase

import (
	"app/ent"
	"app/repository"
	"app/usecase"
	"github.com/stretchr/testify/require"
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

func NewTodoUseCase() *usecase.TodoUseCase {
	r := NewDummyTodoRepository()
	return usecase.NewTodoUseCase(r)
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
		uc := NewTodoUseCase()
		uc.CreateTodo("test1", false, 1)
		uc.CreateTodo("test2", true, 1)

		todos, err := uc.QueryTodosByUserID(1)

		require.Nil(t, err)
		require.Equal(t, 2, len(todos))
		require.Equal(t, "test1", todos[0].Title)
		require.Equal(t, false, todos[0].Done)

		require.Equal(t, "test2", todos[1].Title)
		require.Equal(t, true, todos[1].Done)
	})

	t.Run("Failed to create todo by empty title", func(t *testing.T) {
		uc := NewTodoUseCase()

		todo, err := uc.CreateTodo("", false, 1)

		require.Nil(t, todo)
		require.NotNil(t, err)

		todos, err := uc.QueryTodos(1)

		require.Equal(t, 0, len(todos))
	})

	t.Run("Failed to create todo by invalid userid", func(t *testing.T) {
		uc := NewTodoUseCase()

		todo, err := uc.CreateTodo("test", false, -1)

		require.Nil(t, todo)
		require.NotNil(t, err)

		todos, err := uc.QueryTodos(1)

		require.Equal(t, 0, len(todos))
	})
}
