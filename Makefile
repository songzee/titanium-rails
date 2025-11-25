.PHONY: run build docker-build

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

docker-build:
	docker build -t titanium-rails .

