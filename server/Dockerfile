FROM golang:1.16-alpine

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get -d -v /app
RUN go install -v /app

EXPOSE 8080

ENTRYPOINT go run main.go

# RUN go install github.com/cosmtrek/air@latest
# CMD ["air", "-c", ".air.toml"]