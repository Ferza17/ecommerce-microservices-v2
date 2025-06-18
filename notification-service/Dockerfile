FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/notification-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50053

RUN go build -o notification-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ferza17/ecommerce-v2/notification-service .
CMD [ "./notification-service","migration up" ]
CMD [ "./notification-service","run" ]