FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/payment-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50054

RUN go build -o payment-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ferza17/ecommerce-v2/payment-service .
CMD [ "./payment-service","migration up" ]
CMD [ "./payment-service","run" ]
