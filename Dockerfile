# Use an official Golang runtime as a parent image
FROM golang:1.21.5-alpine3.17

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the local package files to the container
COPY . .

# Build the Go application
RUN go build -o online-home-test .

# Expose the port
EXPOSE 5000

# Run the executable
CMD ["./online-home-test"]