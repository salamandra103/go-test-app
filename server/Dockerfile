FROM golang:1.16-alpine

WORKDIR /app

COPY ./ /app

RUN go mod download

EXPOSE 8080

ENTRYPOINT go run main.go