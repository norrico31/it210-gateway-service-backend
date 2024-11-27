FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:latest

# Install necessary libraries
RUN apk add --no-cache libc6-compat

# Copy the built application from the builder stage
COPY --from=builder /app/main /main

EXPOSE 8083

CMD ["/main"]
