package query

import (
	"app/ent"
	"app/graph"
	"app/graph/generated"
	"app/repository"
	"app/usecase"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
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

func NewTodoRepository() repository.TodoRepository {
	return &dummyTodoRepository{}
}

func NewTodoUseCase() *usecase.TodoUseCase {
	r := NewTodoRepository()
	return usecase.NewTodoUseCase(r)
}

func NewResolver() *graph.Resolver {
	todoUseCase := NewTodoUseCase()
	return &graph.Resolver{TodoUseCase: todoUseCase}
}

func NewClient() *client.Client {
	resolvers := NewResolver()
	config := generated.Config{Resolvers: resolvers}
	schema := generated.NewExecutableSchema(config)
	server := handler.NewDefaultServer(schema)
	return client.New(server)
}

func TestResolver(t *testing.T) {
	t.Run("Add todo query", func(t *testing.T) {
		client := NewClient()

		q := `
			mutation createTodo{ 
				createTodo(input: { title: "todo",  userId: 2 }) {
					id
					title
					done
					userId
				}
			}
		`

		var resp struct {
			CreateTodo struct {
				Id     int
				Title  string
				Done   bool
				UserId int
			}
		}

		client.MustPost(q, &resp)

		require.Equal(t, 1, resp.CreateTodo.Id)
		require.Equal(t, "todo", resp.CreateTodo.Title)
		require.Equal(t, false, resp.CreateTodo.Done)
		require.Equal(t, 2, resp.CreateTodo.UserId)
	})
}
