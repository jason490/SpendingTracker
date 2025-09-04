FROM golang:1.24.3-alpine

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --update gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . /app/


# USER backend

CMD go run ./main.go
