LOCAL_BIN:=$(CURDIR)/bin
NAME=chat
VERSION=1
API_NAME=${NAME}_v${VERSION}

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc@v1.57.0


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/$(API_NAME)
	protoc --proto_path api/$(API_NAME) \
	--go_out=pkg/$(API_NAME) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API_NAME) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/$(API_NAME)/$(NAME).proto

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml
