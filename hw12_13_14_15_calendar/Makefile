VERSION := $(shell git describe --tags | sed -e 's/^v//g' | awk -F "-" '{print $$1}')

RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"
LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_TAG:=1.54.2
PATH:=$(PATH):$(LOCAL_BIN)

# slq migrate
SQL_MIGRATE_CONFIG:=./sql-migrate.yml
SQL_MIGRATE_ENV:=db
.sql-migrate = $(LOCAL_BIN)/sql-migrate $1 -config=$(SQL_MIGRATE_CONFIG) -env=$(SQL_MIGRATE_ENV) $2

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go mod tidy
	GOBIN=$(LOCAL_BIN) go install github.com/google/wire/cmd/wire@v0.5.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.38.0 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.18.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/utrack/clay/v3/cmd/protoc-gen-goclay@v3.1.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/rubenv/sql-migrate/...@v1.5.2

# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: install
install: .bin-deps install-lint

.PHONY: test
test:
	go test -race -coverprofile=coverage.out ./...

.PHONY: show_cover
show_cover:
	go tool cover -html=coverage.out

.PHONY: lint
lint: install-lint
	$(LOCAL_BIN)/golangci-lint run
	
.PHONY: linter-local
linter-local: install-lint
	$(LOCAL_BIN)/golangci-lint run --fix

.PHONY: quality
quality: linter-local tests_integration

.PHONY: clean_cache
clean_cache:
	go clean -cache

.PHONY: build
build:
	go build -ldflags="-X github.com/hound672/otus-hw/hw12_13_14_15_calendar/build.Version=$(VERSION)" -v -o calendar ./cmd

.PHONY: build_docker
build_docker:
	docker build --build-arg VERSION=$(VERSION) -t calendar:$(VERSION) -f docker/Dockerfile .

.PHONY: mockery
mockery:
	$(LOCAL_BIN)/mockery --name $(name) --dir $(dir) --output $(dir)/mocks

.PHONY: wire
wire:
	(cd cmd/application && $(LOCAL_BIN)/wire)

.PHONY: mock
mock:

./PHONY: .sql-migrate
.sql-migrate:
	$(LOCAL_BIN)/sql-migrate status -config=$(SQL_MIGRATE_CONFIG) -env=$(SQL_MIGRATE_ENV)

.PHONY: generate
generate: mock wire .gen-proto

.PHONY: .gen-proto
.gen-proto:
	cd ./pkg/api && \
	protoc -I ./../../api \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=. --go_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-goclay=$(LOCAL_BIN)/protoc-gen-goclay \
		--goclay_out=. \
		--goclay_opt=impl=true,swagger=true,force=false,impl_service_sub_dir=true,impl_type_name_tmpl=Controllers,tests=true,impl_path=../../../internal/controllers,paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway --grpc-gateway_out . \
		--grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	v1/dast.proto

# sql migrations

.PHONY: migrate_status
migrate_status:
	$(call .sql-migrate,status)

# make migrate_new NEW_MIGRATION_NAME
.PHONY: migrate_new
migrate_new:
	$(call .sql-migrate,new,$(RUN_ARGS))

# make migrate_up MIGRATION_ID
# if MIGRATION_ID is empty -> migrate to the latest version
.PHONY: migrate_up
migrate_up:
	$(call .sql-migrate,up,$(RUN_ARGS))

# make migrate_down MIGRATION_ID
# if MIGRATION_ID is empty -> migrate to the first version
.PHONY: migrate_down
migrate_down:
	$(call .sql-migrate,down,$(RUN_ARGS))
