FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/api-gateway

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50000
EXPOSE 3000 3000

RUN go build -o api-gateway

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ferza17/ecommerce-v2/api-gateway .
CMD [ "./api-gateway","run" ]