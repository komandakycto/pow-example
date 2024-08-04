# Use the official Golang image to build the server binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the server binary
RUN make build_verifier

# Use a minimal image to run the server binary
FROM alpine:3.20

WORKDIR /app

# Ensure necessary packages are installed
RUN apk add --no-cache libc6-compat

# Copy the server binary from the builder stage
COPY --from=builder /app/build/verifier /app/verifier

# Copy the quotes file
COPY quotes.txt /app/quotes.txt

# Ensure the binary has executable permissions
RUN chmod +x /app/verifier

# Run the server
ENTRYPOINT ["/app/verifier"]
