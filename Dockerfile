FROM golang:1.23-bookworm AS builder

WORKDIR /src/app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o ./.bin/app ./cmd/go-uber-fx/main.go


EXPOSE 8080
CMD ["./.bin/app"]
