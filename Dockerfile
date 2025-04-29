FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o load-test

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/load-test .

ENTRYPOINT ["./load-test"] 