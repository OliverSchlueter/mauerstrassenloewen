FROM golang:1.24-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main cmd/e2e/main.go

# Create a small final image
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose application port (modify as needed)
EXPOSE 8080

# Run the binary
CMD ["./main"]
