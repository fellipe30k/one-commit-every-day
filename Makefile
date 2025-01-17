# Makefile
default: up

build:
	cp ~/.ssh/id_rsa.pub . && cp ~/.ssh/id_rsa . && docker-compose build && rm id_rsa.pub && rm id_rsa

up:
	docker-compose up

down:
	docker-compose down

test:
	docker-compose run --rm app sh -c "go mod tidy && go run main.go test && git log --oneline"
