FROM golang:1.24.3-alpine AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /go/src/github.com/ferza17/ecommerce-microservices-v2/product-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN GO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -tags musl -o product-service

FROM alpine:latest

# install tzdata for timezones
RUN apk --no-cache add ca-certificates curl tzdata

ENV TZ=Asia/Jakarta
ENV ENV=production

WORKDIR /usr/bin

COPY --from=build /go/src/github.com/ferza17/ecommerce-microservices-v2/product-service/product-service .
COPY . .

RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
