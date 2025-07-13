# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy source code first to resolve all dependencies
COPY . .

# Ensure dependencies are properly resolved and downloaded
RUN go mod tidy && go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cdevents-cli .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1000 cdevents && \
    adduser -D -s /bin/sh -u 1000 -G cdevents cdevents

# Set working directory
WORKDIR /home/cdevents

# Copy binary from builder
COPY --from=builder /app/cdevents-cli .

# Change ownership
RUN chown cdevents:cdevents cdevents-cli

# Switch to non-root user
USER cdevents

# Set entrypoint
ENTRYPOINT ["./cdevents-cli"]
