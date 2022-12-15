.PHONY: setup
setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	brew install grpcui
	export PATH=$PATH:$(go env GOPATH)/bin

.PHONY: clean
clean:
	rm -rf server/pb/
	rm -rf client/pb/

.PHONY: gen
gen: clean
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

.PHONY: run-server
run-server:
	cd server && go run main.go

.PHONY: run-client
run-client:
	cd client && go run main.go


.PHONY: build-server
build-server:
	cd server && go build -o ../bin/server  main.go

.PHONY: build-client
build-client:
	cd client && go build -o ../bin/client  main.go

.PHONY: build
build: build-server build-client


.PHONY: build-server-image
build-server-image:
	cd server && docker build . -t demoserver

.PHONY: run-server-image
run-server-image: build-server-image
	docker run -it --rm -p 8080:8080 demoserver

.PHONY: test
test:
	echo "Test missing!"

