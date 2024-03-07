# Use the official Golang image as a base
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Use a minimal base image to keep the image size small
FROM alpine:latest

# Create a non-root user with UID 1001
RUN adduser -u 1001 -D appuser

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/app .

# Change the ownership of the binary to the non-root user
RUN chown appuser:appuser /app/app

# Tell Docker that all future commands should run as the appuser user
USER appuser

# Command to run the executable
CMD ["./app"]
