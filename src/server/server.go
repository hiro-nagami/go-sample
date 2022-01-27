package server

import (
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
