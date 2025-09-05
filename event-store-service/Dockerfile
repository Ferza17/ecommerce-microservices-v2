FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-microservices-v2/event-store-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 50052

RUN CGO_ENABLED=0 GOOS=linux go build -o event-store-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

ENV ENV "production"
WORKDIR /usr/bin

COPY --from=build /go/src/github.com/ferza17/ecommerce-microservices-v2/event-store-service/event-store-service .
COPY . .

RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]