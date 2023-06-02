FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY src/ ./src/

RUN go build -o ./bin/book-service-app ./src/

EXPOSE 3000

CMD ["./bin/book-service-app"]