# ---- Build Stage ----
FROM golang:1.25-alpine AS builder

# Install git + CA certificates for secure HTTPS
RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o url-shortener ./


# ---- Runtime Stage ----
FROM alpine:latest

# Install CA certs in runtime too (for HTTPS APIs, if needed)
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/url-shortener .

EXPOSE 5000

CMD ["./url-shortener"]
