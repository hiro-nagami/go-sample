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
	"log"
	"net/http"
)

const defaultPort = "8080"

type Server struct {
	router  *chi.Mux
	handler http.Handler
	path    string
	port    string
}

func (s *Server) Serve() {
	s.router.Handle(s.path, s.handler)
	log.Fatal(http.ListenAndServe(":"+s.port, s.router))
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
		router:  router,
		handler: handler,
		path:    "/graph",
		port:    defaultPort,
	}
}
