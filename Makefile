# Makefile
default: up

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

test:
	docker-compose run --rm app sh -c "go mod tidy && go run main.go test && git log --oneline"