ARG GOLANG_IMAGE

FROM ${GOLANG_IMAGE}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN /go/bin/swag init -g ./cmd/main.go
RUN echo $GOPATH && go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]