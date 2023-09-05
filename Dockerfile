FROM    golang:1.21.0-bullseye as builder
RUN     mkdir "/app"
WORKDIR "/app"
COPY    . /app/
RUN     go mod download
RUN     go build -o todo-app ./cmd/main.go
CMD     ["./todo-app"]
