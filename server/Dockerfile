# FROM golang:1.16-alpine

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY *.go ./

# EXPOSE 8080



FROM golang:1.16-alpine

WORKDIR /app

COPY ./ /app

RUN go mod download

EXPOSE 8080

ENTRYPOINT go run main.go