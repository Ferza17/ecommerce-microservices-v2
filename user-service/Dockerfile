FROM golang:1.24.3-alpine AS build

WORKDIR /go/src/github.com/ferza17/ecommerce-v2/user-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 50056

RUN go build -o user-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ferza17/ecommerce-v2/user-service .
CMD [ "./user-service","migration up" ]
CMD [ "./user-service","run" ]