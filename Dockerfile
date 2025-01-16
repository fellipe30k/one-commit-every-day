# Dockerfile
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y git && go mod tidy && go build -buildvcs=false -o /bin/one-commit-every-day

FROM debian:bookworm
WORKDIR /app
RUN apt-get update && apt-get install -y git wget && \
    wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz && \
    rm go1.20.7.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/bin/go
COPY --from=builder /bin/one-commit-every-day /usr/local/bin/
COPY . /app
COPY .git /app/.git
CMD ["one-commit-every-day"]
