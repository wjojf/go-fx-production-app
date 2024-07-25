FROM golang:1.23-rc-bookworm AS builder

# ARG GOOSE_DRIVER
# ARG GOOSE_DBSTRING
# ARG GOOSE_MIGRATION_DIR


# RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

WORKDIR /src/app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o ./.bin/app ./cmd/go-uber-fx/main.go

# RUN ["make", "up"]

EXPOSE 8080
CMD ["./.bin/app"]
