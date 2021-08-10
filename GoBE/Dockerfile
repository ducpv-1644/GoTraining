FROM golang:1.11.0-alpine

# go get uses git to fetch modules
RUN apk add --no-cache git

RUN mkdir -p /go/src/gobe
WORKDIR /go/src/gobe

CMD go run main.go
