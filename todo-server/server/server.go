package server

import (
	"app/repository"
	"app/usecase"
	"net"
)

type Server interface {
	Serve(lis net.Listener) error
	Inject(services *Services)
}

type Services struct {
	Todo *usecase.TodoUseCase
	User *usecase.UserUseCase
}

func (s *Services) InjectTodo(todo *usecase.TodoUseCase) {
	s.Todo = todo
}

func (s *Services) InjectUser(user *usecase.UserUseCase) {
	s.User = user
}

func NewService() *Services {
	services := Services{}

	todo := &usecase.TodoUseCase{
		Repo: repository.NewTodoRepository(),
	}

	services.InjectTodo(todo)

	return &services
}