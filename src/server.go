package main

import (
	"app/repository"
	sv "app/server"
	"app/server/grpc"
	"app/usecase"
)

func main() {
	services := &sv.Services{}

	todo := &usecase.TodoUseCase{
		Repo: repository.NewTodoRepository(),
	}

	services.Inject(todo)

	var server sv.Server = grpc.NewServer()
	//var server sv.Server = graphql.NewServer()
	server.Inject(services)
	server.Serve()
}
