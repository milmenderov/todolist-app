FROM golang:1.21-bullseye

WORKDIR /app

COPY ./ ./


RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]