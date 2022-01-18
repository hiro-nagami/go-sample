package main

import (
    _ "fmt"
    "log"
    "net/http"
    "os"

    "app/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "0.0.0.0:8081/graph"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func mustGet(arg string) string {
	env := os.Getenv(arg)
	if env == "" {
		panic("env not found")
	}
	return env
}
