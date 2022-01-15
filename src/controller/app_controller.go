package controller

import (
	"app/graph"
	"app/graph/generated"
	"app/repository"
	"app/usecase"
	"app/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const defaultPort = "8080"

type Server struct {
	handler http.Handler
	path    string
}

func NewGraphqlServer() *Server {
	todoRepository := repository.NewTodoRepository()

	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	resolver := &graph.Resolver{TodoUseCase: todoUseCase}
	config := generated.Config{Resolvers: resolver}
	schema := generated.NewExecutableSchema(config)

	handler := handler.NewDefaultServer(schema)
	return &Server{
		handler: handler,
		path:    "/graph",
	}
}

type AppController interface {
	Inject(server *Server)
	Serve()
}

type appController struct {
	router *chi.Mux
	server *Server
	port   string
}

func NewAppController() AppController {
	port := utils.MustGet("PORT")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.New(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	return &appController{
		router: router,
		server: nil,
		port:   port,
	}
}

func (controller *appController) Inject(server *Server) {
	controller.server = server
}

func (controller *appController) Serve() {
	controller.router.Handle(controller.server.path, controller.server.handler)
	log.Fatal(http.ListenAndServe(":"+controller.port, controller.router))
}
