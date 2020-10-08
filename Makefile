SHELL=/bin/bash

.PHONY: all
all: deps gen lint build upd system_test down

.PHONY: ci
ci: all

.PHONY: deps
deps:
	go mod tidy

.PHONY: system_test
system_test:
	go test ./system_test/. -v --godog.format=pretty

.PHONY: gen
gen:
	protoc -I api api/service.proto --go_out=plugins=grpc:api

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -o artifacts/svc .

.PHONY: up
up:
	docker-compose -f docker-composition/default.yml -f docker-composition/build-mask.yml up --build --remove-orphans

.PHONY: upd
upd:
	docker-compose -f docker-composition/default.yml -f docker-composition/build-mask.yml up -d --build --remove-orphans

.PHONY: down
down:
	docker-compose -f docker-composition/default.yml -f docker-composition/build-mask.yml down

.PHONY: lint
lint:
	golangci-lint run
