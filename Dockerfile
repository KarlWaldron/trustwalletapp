# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o blockchain-client .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/blockchain-client .
EXPOSE 8080
CMD ["./blockchain-client"]
