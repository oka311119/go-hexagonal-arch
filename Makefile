.PHONY: build up down mockgen lint tidy test

build:
	docker-compose build

up:
	docker-compose up --build

down:
	docker-compose down -v

// TODO: mockgenの実行を自動化する
mockgen:
	./script/generate_mocks.sh

lint:
	errcheck ./...
	staticcheck ./...
	goimports -w ./
	go vet ./...

tidy:
	go mod tidy

test:
	go test -v ./...
