FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/payment-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50055

RUN go build -o payment-service