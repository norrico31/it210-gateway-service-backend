FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Final production image with Nginx
FROM alpine:latest

# Install Nginx and necessary libraries
RUN apk add --no-cache nginx libc6-compat

COPY nginx/nginx.conf /etc/nginx/nginx.conf

# Copy the built application from the builder stage
COPY --from=builder /app/main /main

EXPOSE 80 8083

CMD ["sh", "-c", "nginx -g 'daemon off;' & /main"]
