FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/product-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50055

RUN go build -o product-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ferza17/ecommerce-v2/product-service .
CMD [ "./product-service","migration up" ]
CMD [ "./product-service","migration elasticsearch" ]
CMD [ "./product-service","run" ]