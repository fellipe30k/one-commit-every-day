# Makefile
default: up

up:
	docker-compose up --build

down:
	docker-compose down

test:
	docker-compose run --rm app sh -c "go run main.go && git log --oneline"