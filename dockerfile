FROM golang:1.19

RUN mkdir /build

WORKDIR /build

COPY . .

RUN go get -u github.com/gin-gonic/gin

EXPOSE 3000

CMD ["go", "run", "main.go"]