# gRPC on Go - Dockerized ![go](https://img.shields.io/badge/Go-00ADD8?style=plastic&logo=go&logoColor=white) ![docker](https://img.shields.io/badge/Docker-2CA5E0?style=plastic&logo=docker&logoColor=white)

This projects acts as a starting point for building a server(in Go) & a client(in Go) that communicates via gRPC.

Blog for the same coming soon!

## Dev Setup 🏗

```bash
make setup
```
will run the following:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
brew install protobuf
brew install clang-format
brew install grpcurl
brew install grpcui
export PATH=$PATH:$(go env GOPATH)/bin

```

## Invoking RPCs 🚀

```bash
# Note: since we are not using TLS all the calls are with -plaintext flag
grpcui -plaintext localhost:8080 # Use grpc-ui to use UI
```

```bash
# Note: since we are not using TLS all the calls are with -plaintext flag
grpcurl -plaintext localhost:8080 list # introspect the service
grpcurl -plaintext localhost:8080 <service-name>.<rpc name> # to get rpc output
```

## Useful make commands 👨🏻‍💻

Generate Go stubs

```bash
make gen
```

Run server

```bash
make run-server
```

Run client

```bash
make run-client
```

Build Go Binaries

```bash
make build # To build both
```
```bash
make build-server
```
```bash
make build-client
```

Build server image and run

```bash
make build-server-image
```
```bash
make run-server-image # Build and run server image
```

## Demo 🔥

![render1671107836080](https://user-images.githubusercontent.com/54388005/207862391-1d339742-ee5b-44ff-83fb-82ef520a8d28.gif)


![render1671108620283](https://user-images.githubusercontent.com/54388005/207863536-61e4d173-0c84-4536-bade-2540b698b04b.gif)
