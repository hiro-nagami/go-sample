package server

import (
	"app/usecase"
)

type Server interface {
	Serve()
	Inject(services *Services)
}

type Services struct {
	Todo *usecase.TodoUseCase
}

func (s *Services) Inject(todo *usecase.TodoUseCase) {
	s.Todo = todo
}
