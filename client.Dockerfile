# Use the official Golang image to build the client binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the client binary
RUN make build_prover

# Use a minimal image to run the client binary
FROM alpine:3.20

WORKDIR /app

# Ensure necessary packages are installed
RUN apk add --no-cache libc6-compat

# Copy the client binary from the builder stage
COPY --from=builder /app/build/prover /app/prover

# Ensure the binary has executable permissions
RUN chmod +x /app/prover

# Run the client
ENTRYPOINT ["/app/prover"]
