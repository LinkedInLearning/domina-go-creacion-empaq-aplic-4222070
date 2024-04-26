# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pkm .

# Stage 2: Create the final image
FROM scratch

WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/pkm .

# the tls certificates:
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set the entry point for the container
ENTRYPOINT ["./pkm"]
