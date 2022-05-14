package main

import (
	sv "app/server"
	"app/server/grpc"
)

func main() {
	server := grpc.NewServer()
	server.Inject(sv.NewService())
	server.Serve(nil)
}
