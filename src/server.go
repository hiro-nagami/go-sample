package main

import "app/controller"

func main() {
	app := controller.NewAppController()
	server := controller.NewGraphqlServer()
	app.Inject(server)
	app.Serve()
}
