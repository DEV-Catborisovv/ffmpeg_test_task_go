FROM golang:1.22.0 AS builder
WORKDIR /go/delivery
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o main cmd/app/main.go

CMD ["./main"]