# Stage 1: Build the Go binary
FROM golang:latest as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o backend cmd/app/main.go

# Stage 2: Run the Go binary
FROM debian:bullseye-slim

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Copy any other necessary files
COPY config.yaml .

# Expose the port that the application will run on
EXPOSE 3000

# Command to run the binary
CMD ["./myapp"]
