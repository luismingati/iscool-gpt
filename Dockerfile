# Stage 1: Builder
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies (cached layer)
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0: Disable CGO for static binary (required for scratch)
# -ldflags="-w -s": Strip debug information to reduce binary size
# -a: Force rebuild of all packages
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o iscool-gpt .

# Stage 2: Final image using scratch
FROM scratch

# Copy SSL certificates from builder (required for HTTPS calls to Gemini API)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder
COPY --from=builder /app/iscool-gpt /iscool-gpt

# Expose port 8080
EXPOSE 8080

# Run the application
ENTRYPOINT ["/iscool-gpt"]
