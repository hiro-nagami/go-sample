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
}

func (s *Services) Inject(todo *usecase.TodoUseCase) {
	s.Todo = todo
}
