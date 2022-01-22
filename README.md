## Glang sample development
1. Setup gRPC
```shell
$ brew install protobuf
$ go get -u google.golang.org/grpc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
 ```

2. Generate gRPC files
```shell
$  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false proto/*.proto
```

### Setup docker
1. Run `script/setup.sh`
2. Run `go generate ./ent`
3. Run `go run github.com/99designs/gqlgen`
4. Run `go get -u google.golang.org/grpc`
5. Run `docker-compose up`


### Testing
1. Run `docker-compose run --rm server go test -v app/test/...`