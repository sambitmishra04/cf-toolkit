# 1. Build Stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# 1. Build Stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o cf-toolkit .

# 2. Run Stage (Tiny image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cf-toolkit .
# Note: We use environment variables for credentials in production.

CMD ["./cf-toolkit"]