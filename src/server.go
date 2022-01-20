package main

import (
	"app/repository"
	sv "app/server"
	"app/server/grpc"
	"app/usecase"
)

func NewService() *sv.Services {
	services := sv.Services{}

	todo := &usecase.TodoUseCase{
		Repo: repository.NewTodoRepository(),
	}

	services.Inject(todo)

	return &services
}

func NewServer() sv.Server {
	return grpc.NewServer()
}

func main() {
	server := NewServer()
	server.Inject(NewService())
	server.Serve(nil)
}
