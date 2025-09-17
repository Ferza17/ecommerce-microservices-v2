FROM golang:1.25.0-alpine AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /go/src/github.com/ferza17/ecommerce-microservices-v2/user-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN GO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -tags musl -o user-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl tzdata

ENV TZ=Asia/Jakarta
ENV ENV "production"
WORKDIR /usr/bin

COPY --from=build /go/src/github.com/ferza17/ecommerce-microservices-v2/user-service/user-service .
COPY . .
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]