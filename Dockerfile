# Build stage
FROM golang:1.23-alpine AS builder

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o mcp-scan .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/mcp-scan .

# Copy demo configurations
COPY --from=builder /app/demos ./demos

# Create a non-root user
RUN adduser -D -s /bin/sh mcpuser
USER mcpuser

# Expose no ports (this is a CLI tool)

# Set entrypoint
ENTRYPOINT ["./mcp-scan"]

# Default to showing help
CMD ["--help"]

# Metadata
LABEL org.opencontainers.image.title="MCP Scan"
LABEL org.opencontainers.image.description="Security scanner for Model Context Protocol (MCP) configurations"
LABEL org.opencontainers.image.source="https://github.com/AndreaGriffiths11/mcp-config-scan"
LABEL org.opencontainers.image.licenses="MIT"