# Build stage
FROM golang:1.21-alpine AS builder

# Build arguments
ARG VERSION=dev
ARG REVISION=unknown
ARG BUILDTIME=unknown

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go.mod and go.sum for better layer caching
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Ensure dependencies are properly resolved
RUN go mod tidy

# Build the binary with version information
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -ldflags "-s -w -X main.version=${VERSION} -X main.revision=${REVISION} -X main.buildTime=${BUILDTIME}" \
    -o cdevents-cli .

# Final stage
FROM alpine:latest

# Build arguments for metadata
ARG VERSION=dev
ARG REVISION=unknown
ARG BUILDTIME=unknown

# Add metadata labels
LABEL org.opencontainers.image.title="CDEvents CLI" \
      org.opencontainers.image.description="Command-line tool for generating and sending CDEvents" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.revision="${REVISION}" \
      org.opencontainers.image.created="${BUILDTIME}" \
      org.opencontainers.image.source="https://github.com/brunseba/cdevents-tools" \
      org.opencontainers.image.url="https://github.com/brunseba/cdevents-tools" \
      org.opencontainers.image.documentation="https://github.com/brunseba/cdevents-tools/blob/main/README.md" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.vendor="CDEvents Community"

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
