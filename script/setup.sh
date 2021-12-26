#!/bin/sh

docker-compose build
docker-compose run --rm server go mod tidy
docker-compose run --rm server go get github.com/99designs/gqlgen
docker-compose run --rm server go run github.com/99designs/gqlgen init
docker-compose run --rm graphql-playground-server go mod tidy