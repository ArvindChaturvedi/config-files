# Use the official golang image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o myapp

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]