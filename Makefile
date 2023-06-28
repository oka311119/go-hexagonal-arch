.PHONY: build up down mockgen test

build:
	docker-compose build

up:
	docker-compose up --build

down:
	docker-compose down -v

// TODO: mockgenの実行を自動化する
mockgen:
	./script/generate_mocks.sh

test:
	go test -v ./...
