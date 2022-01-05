## Glang sample development

### Setup docker
1. Run `script/setup.sh`
2. Run `go generate ./ent`
3. Run `go run github.com/99designs/gqlgen`
4. Run `docker-compose up`


### Testing
1. Run `docker-compose run --rm server go test -v ./test`