FROM golang:1.24.6-alpine

WORKDIR /

COPY . .

RUN go mod download

RUN go run .

EXPOSE 8080
