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
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	go install github.com/bufbuild/buf/cmd/buf && \
	go install github.com/vektra/mockery/v2 && \
	go install github.com/pressly/goose/v3/cmd/goose

MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: install-lint
install-lint:
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2

.PHONY: lint
lint: install-lint
	$(info Running lint...)
	$(LOCAL_BIN)/golangci-lint run --config=.golangci.yml ./...

.PHONY: test
test:
	$(info Running tests...)
	go test ./...

.PHONY: integration-test
integration-test:
	$(info Running tests...)
	# TODO take goose from setup-ci job
	go install github.com/pressly/goose/v3/cmd/goose
	./migrate.sh test
	go test -tags=integration ./...

.PHONY: setup-ci
setup-ci:
	export GOBIN="/usr/local/go/bin" && \
	go env && \
	make .deps && \
	buf build && \
	buf generate && \
	go generate ./...