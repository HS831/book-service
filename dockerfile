FROM golang:1.19

RUN mkdir /build

WORKDIR /build

COPY . .

RUN go mod vendor

EXPOSE 3000

CMD ["go", "run", "main.go"]