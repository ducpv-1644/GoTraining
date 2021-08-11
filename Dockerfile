FROM golang:1.15.7-alpine

# go get uses git to fetch modules
RUN apk add --no-cache git

RUN mkdir -p /go/src/gobe
WORKDIR /go/src/gobe
COPY go.mod go.sum ./
RUN go mod download
COPY . .

CMD go run app/main.go
