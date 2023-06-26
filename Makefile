.PHONY: run test

run:
	docker-compose up --build
	# go run main.go

test:
	go test -v ./...
