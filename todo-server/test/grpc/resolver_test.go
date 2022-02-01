package grpc

import (
	"app/ent"
	pb "app/proto"
	"app/repository"
	sv "app/server"
	grpcServer "app/server/grpc"
	"app/usecase"
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	server := NewServer()
	server.Inject(NewService())

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

type dummyTodoRepository struct {
	todos []*ent.Todo
	count int
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
	return &dummyTodoRepository{
		todos: []*ent.Todo{},
		count: 0,
	}
}

type dummyUserRepository struct {
	users []*ent.User
	count int
}

func (repo *dummyUserRepository) CreateUser(name string, sex int) (*ent.User, error) {
	repo.count++

	user := &ent.User{
		ID:   repo.count,
		Name: name,
		Sex:  sex,
	}
	repo.users = append(repo.users, user)

	return user, nil
}

func (repo *dummyUserRepository) QueryUsers(id int) ([]*ent.User, error) {
	users := []*ent.User{}

	for i := range repo.users {
		if repo.users[i].ID == id {
			users = append(users, repo.users[i])
		}
	}

	return users, nil
}

func NewUserRepository() repository.UserRepository {
	return &dummyUserRepository{
		users: []*ent.User{},
		count: 0,
	}
}

func NewService() *sv.Services {
	services := sv.Services{}

	todo := &usecase.TodoUseCase{
		Repo: NewTodoRepository(),
	}
	user := &usecase.UserUseCase{
		Repo: NewUserRepository(),
	}

	services.InjectTodo(todo)
	services.InjectUser(user)

	return &services
}

func NewServer() sv.Server {
	return grpcServer.NewServer()
}

func TestPostgRPCMessasge(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	require.Nil(t, err)

	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	newTodo := &pb.NewTodo{Title: "test1", Done: false, UserId: 1}

	resp, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{Todo: newTodo})
	require.Nil(t, err)

	require.Equal(t, resp.GetTodo().Id, int32(1))
	require.Equal(t, resp.GetTodo().Title, "test1")
	require.Equal(t, resp.GetTodo().Done, false)
	require.Equal(t, resp.GetTodo().UserId, int32(1))

}
