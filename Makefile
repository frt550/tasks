run-tg:
	go run cmd/bot/main.go
run-task:
	go run cmd/task/grpc/main.go & go run cmd/task/rest/main.go
run-backup:
	go run cmd/backup/grpc/main.go & go run cmd/backup/rest/main.go

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql