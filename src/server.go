package main

import (
	"app/server"
	"app/server/grpc"
)

func main() {
	var server server.Server
	//server = graphql.NewServer()
	server = grpc.NewServer()

	server.Serve()
}
