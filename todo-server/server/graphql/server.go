package graphql

import (
	"app/graph"
	"app/graph/generated"
	"app/repository"
	"app/server"
	"app/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"net"
	"net/http"
)

const defaultPort = "8080"

type Server struct {
	Router   *chi.Mux
	Handler  http.Handler
	Path     string
	Port     string
	Services *server.Services
}

func (s *Server) Inject(services *server.Services) {
	s.Services = services
}

type Services struct {
	Todo *usecase.TodoUseCase
}

func (s *Server) Serve(lis net.Listener) error {
	s.Router.Handle(s.Path, s.Handler)
	return http.ListenAndServe(":"+s.Port, s.Router)
}

func (s *Services) Inject(todo *usecase.TodoUseCase) {
	s.Todo = todo
}

func NewRouter() *chi.Mux {
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

	return router
}

func NewHandler(resolver *graph.Resolver) http.Handler {
	config := generated.Config{Resolvers: resolver}
	schema := generated.NewExecutableSchema(config)

	return handler.NewDefaultServer(schema)
}

func NewServer() server.Server {
	todoRepository := repository.NewTodoRepository()

	todoUseCase := usecase.NewTodoUseCase(todoRepository)
	resolver := &graph.Resolver{TodoUseCase: todoUseCase}

	handler := NewHandler(resolver)
	router := NewRouter()

	return &Server{
		Router:   router,
		Handler:  handler,
		Path:     "/graph",
		Port:     defaultPort,
		Services: nil,
	}
}
