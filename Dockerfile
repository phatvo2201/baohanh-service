# Dockerfile
FROM golang:1.24.1-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app
RUN go build -o /app/main ./cmd/server

# Expose port
EXPOSE 8080

# Run the built binary
CMD ["/app/main"]
