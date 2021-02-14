.PHONY: lint run build binary test all

VERSION = $(shell git rev-parse --verify HEAD)

all: lint run

build: binary build

binary:
	go build -ldflags "-X main.version=$(VERSION)" cmd/slidups.go

build:
	docker build --build-arg version=$(VERSION) .

run:
	go run cmd/slidups.go --upload.destination=/tmp

lint:
	@go get -u golang.org/x/lint/golint
	golint ./...

test:
	go test -v -cover ./...
