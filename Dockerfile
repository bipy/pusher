# Build stage
FROM golang:1.20-alpine AS builder

ARG VERSION=dev
ARG BUILD_DATE
ARG VCS_REF

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with version information
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE} -X main.GitCommit=${VCS_REF}" \
    -o pusher .

# Runtime stage
FROM alpine:latest

ARG VERSION=dev
ARG BUILD_DATE
ARG VCS_REF

# Add metadata labels
LABEL org.opencontainers.image.title="Pusher"
LABEL org.opencontainers.image.description="A lightweight, secure Telegram message delivery API"
LABEL org.opencontainers.image.authors="bipy <notbipy@gmail.com>"
LABEL org.opencontainers.image.url="https://github.com/bipy/pusher"
LABEL org.opencontainers.image.source="https://github.com/bipy/pusher"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.created="${BUILD_DATE}"
LABEL org.opencontainers.image.revision="${VCS_REF}"
LABEL org.opencontainers.image.licenses="MIT"

WORKDIR /app

# Install ca-certificates for HTTPS (critical for Telegram API)
# Create non-root user and group (may already exist, so ignore errors)
RUN apk --no-cache add ca-certificates && \
    addgroup -g 1000 pusher 2>/dev/null || true && \
    adduser -D -u 1000 -G pusher pusher 2>/dev/null || true

# Copy binary and entrypoint from builder
COPY --from=builder --chown=pusher:pusher /build/pusher .
COPY --from=builder --chown=pusher:pusher /build/entrypoint.sh .

# Make entrypoint executable
RUN chmod +x entrypoint.sh

# Environment variables with defaults
ENV SERVER_HOST="0.0.0.0" \
    SERVER_PORT=3333 \
    TG_TOKEN="" \
    CHAT_ID="" \
    SECURE_KEY=""

# Switch to non-root user
USER pusher

EXPOSE $SERVER_PORT

# Note: healthcheck can be added via docker-compose or kubernetes if needed
# Example: healthcheck with wget (smaller than curl)
# HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
#     CMD wget --quiet --tries=1 --spider http://localhost:${SERVER_PORT}/pulse || exit 1

ENTRYPOINT ["/app/entrypoint.sh"]