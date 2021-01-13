FROM golang:alpine as builder

WORKDIR /github.com/hieronimusbudi/go-fiber-bookstore-item-api
COPY go.mod go.sum ./
RUN go mod download && go get github.com/codegangsta/gin
COPY . .

CMD gin --immediate -a 9000 -p 9010 run server.go