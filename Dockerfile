FROM golang:1.22.0-alpine AS builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev sqlite
COPY . .
RUN go build -o app ./cmd/main.go
CMD ["./app"]