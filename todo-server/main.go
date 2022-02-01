package main

import (
	sv "app/server"
)

func main() {
	server := sv.NewServer()
	server.Inject(sv.NewService())
	server.Serve(nil)
}
