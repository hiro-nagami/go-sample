package grpc

import (
	pb "app/proto"
	sv "app/server"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const defaultPort = "8080"

type Server struct {
	Server *grpc.Server
	Path   string
	Port   string
	Helper *ServicesHelper
}

type ServicesHelper struct {
	services *sv.Services
}

func (s *Server) Inject(services *sv.Services) {
	s.Helper = &ServicesHelper{
		services: services,
	}
}

func (h *ServicesHelper) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.TodoResponse, error) {
	newTodo, err := h.services.Todo.CreateTodo(request.Todo.Title, request.Todo.Done, int(request.Todo.UserId))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return &pb.TodoResponse{
		Todo: &pb.Todo{
			Id:     int32(newTodo.ID),
			Title:  newTodo.Title,
			Done:   newTodo.Done,
			UserId: int32(newTodo.UserID),
		},
	}, nil
}

func (h *ServicesHelper) Todos(ctx context.Context, request *pb.TodosRequest) (*pb.TodosResponse, error) {
	id := int(request.UserId)
	fetched, err := h.services.Todo.QueryTodosByUserID(id)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	todos := []*pb.Todo{}

	for _, t := range fetched {
		todos = append(todos, &pb.Todo{
			Id:     int32(t.ID),
			Title:  t.Title,
			Done:   t.Done,
			UserId: int32(t.UserID),
		})
	}

	return &pb.TodosResponse{
		Todos: todos,
	}, nil
}

func (s *Server) Serve(lis net.Listener) {
	var err error = nil

	if lis == nil {
		lis, err = net.Listen("tcp", ":"+defaultPort)
	}

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterTodoServiceServer(s.Server, s.Helper)
	reflection.Register(s.Server)

	log.Fatal(s.Server.Serve(lis))
}

func NewServer() sv.Server {
	server := grpc.NewServer()

	return &Server{
		Server: server,
		Path:   "/grpc",
		Port:   defaultPort,
		Helper: nil,
	}
}
