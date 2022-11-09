FROM golang:1.19-buster AS build

ARG MIGRATE_VERSION=4.7.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
RUN go mod download

COPY ./ ./
RUN go build -o /my-server

COPY /entrypoint.sh /usr/local/bin
RUN ln -s usr/local/bin/entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE "9000"
