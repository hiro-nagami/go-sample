#!/bin/sh

docker-compose build
docker-compose run --rm server go mod tidy
docker-compose run --rm server go get github.com/99designs/gqlgen
docker-compose run --rm server go run github.com/99designs/gqlgen init
docker-compose run --rm server go get github.com/99designs/gqlgen/cmd@v0.14.0
docker-compose run --rm server go run github.com/99designs/gqlgen
docker-compose run --rm server go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
docker-compose run --rm server go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
docker-compose run --rm graphql-playground-server go mod tidy