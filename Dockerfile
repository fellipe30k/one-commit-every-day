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
RUN git config --global user.name "Your Name" && \
    git config --global user.email "your.email@example.com"
COPY . /app
COPY .git /app/.git
# Prepara o diretório SSH e copia a chave SSH
RUN mkdir -p /root/.ssh
COPY id_rsa.pub /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa.pub && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts
CMD ["one-commit-every-day"]
