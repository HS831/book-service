FROM golang:1.19

WORKDIR /book-service

COPY go.mod go.sum ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8000

CMD ["/docker-gs-ping"]