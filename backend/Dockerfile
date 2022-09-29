# Use the official Go image as the base image
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the source code
COPY backend/ .

# Build the Go application
RUN go build -o gaizeAPI .

# Use a minimal base image to run the compiled binary
FROM gcr.io/distroless/base

# Copy the binary from the builder image
COPY --from=builder /app/gaizeAPI /gaizeAPI

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["/gaizeAPI"]
