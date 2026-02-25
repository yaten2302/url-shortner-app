# Build stage
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o url-shortener ./

# Runtime stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/url-shortener .

EXPOSE 5000

CMD ["./url-shortener"]
