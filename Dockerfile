# Use the official Golang image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Install git and necessary build tools
RUN apk add --no-cache git gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/web

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 