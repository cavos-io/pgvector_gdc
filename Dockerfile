# Use the official Golang image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose the application port
EXPOSE 8080

# Run the app
CMD ["./main"]
