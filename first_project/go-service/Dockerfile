# Use the official Go image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first
COPY go.mod ./


# Download dependencies (will cache this layer if no changes)
RUN go mod download

# Copy the rest of the Go application
COPY . .

# Build the Go application
RUN go build -o main .

# Run the Go application
CMD ["./main"]