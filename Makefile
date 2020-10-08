SHELL=/bin/bash

.PHONY: all
all: deps gen lint build upd test down

.PHONY: ci
ci: all

.PHONY: deps
deps:
	go mod tidy

.PHONY: test
test:
	@go get -u github.com/cucumber/godog/cmd/godog
	(cd system_test && exec godog --strict;)

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
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --vendor --install --force
	gometalinter ./... --concurrency=1 --vendor --skip=vendor --exclude=\.*mock\.*\.go --exclude=\.*test\.*\.go --exclude=vendor\.* --cyclo-over=15 --deadline=10m --disable-all \
	--enable=golint \
	--enable=errcheck \
	--enable=vet \
	--enable=deadcode \
	--enable=gocyclo \
	--enable=varcheck \
	--enable=structcheck \
	--enable=vetshadow \
	--enable=ineffassign \
	--enable=interfacer \
	--enable=unconvert \
	--enable=goconst \
	--enable=gosimple \
	--enable=staticcheck \
	--enable=gosec \
	--exclude="\bexported .* should have comment .*or be unexported\b"
