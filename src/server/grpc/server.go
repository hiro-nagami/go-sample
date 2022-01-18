package grpc

import (
	pb "app/proto"
	"app/server"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const defaultPort = "8080"

type Server struct {
	server  *grpc.Server
	handler pb.GreeterServer
	path    string
	port    string
}

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello gRPC, %s.\n", req.Name),
	}, nil
}

func (s *Server) Serve() {
	lis, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterGreeterServer(s.server, s.handler)
	reflection.Register(s.server)

	log.Fatal(s.server.Serve(lis))
}

func NewServer() server.Server {
	server := grpc.NewServer()
	handler := &GreeterService{}

	return &Server{
		server:  server,
		handler: handler,
		path:    "/grpc",
		port:    defaultPort,
	}
}
