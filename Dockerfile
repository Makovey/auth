FROM golang:1.22.2-alpine AS builder

RUN apk add --upgrade --no-cache ca-certificates && update-ca-certificates

COPY . /github.com/Makovey/microservices/auth/source
WORKDIR /github.com/Makovey/microservices/auth/source

RUN go mod download
RUN go build -o ./bin/auth_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Makovey/microservices/auth/source/bin/auth_server .

CMD ["./auth_server"]