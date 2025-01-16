# Dockerfile
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y git && go mod tidy && go build -o /bin/one-commit-every-day

FROM debian:bullseye
WORKDIR /app
RUN apt-get update && apt-get install -y git golang
COPY --from=builder /bin/um-commit-todos-os-dias /usr/local/bin/
COPY . /app
COPY .git /app/.git
CMD ["one-commit-every-day"]