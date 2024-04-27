# Start from a base image containing the Go runtime
FROM golang:1.17-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from the Alpine Linux base image
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the static files (HTML, CSS, JavaScript, etc.) into the container
COPY assets/ ./assets/

# Expose the port on which the Go application will run
EXPOSE 9091

# Command to run the executable
CMD ["./main"]
