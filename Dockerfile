FROM golang:1.24.6-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

CMD [ "go", "run", "." ]
