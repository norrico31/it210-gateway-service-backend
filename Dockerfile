FROM golang:1.23-alpine AS builder

WORKDIR /app/gateway/

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache nginx libc6-compat

COPY nginx/nginx.conf /etc/nginx/nginx.conf

COPY --from=builder /app/gateway/main /main

EXPOSE 80 8080

CMD ["sh", "-c", "nginx && /main"]
