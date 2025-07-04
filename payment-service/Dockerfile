FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-microservices-v2/payment-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o payment-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV ENV "production"
WORKDIR /usr/bin

COPY --from=build /go/src/github.com/ferza17/ecommerce-microservices-v2/payment-service/payment-service .
COPY . .
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]