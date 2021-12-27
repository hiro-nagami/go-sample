package main

import (
    "log"
    "net/http"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
    "app/graph/generated"
	"app/graph"
    "app/utils"
)


const defaultPort = "8080"

func main() {

	port := utils.MustGet("PORT")

    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Use(cors.New(cors.Options{
        // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
        AllowedOrigins:   []string{"https://*", "http://*"},
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    }).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/query", srv)

    log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}