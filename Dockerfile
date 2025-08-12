# Stage 1: Build
FROM golang:alpine AS builder

# Install necessary packages
RUN apk add --no-cache git build-base

# Chromium dependencies (opsional di build stage)
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    libc6-compat

# Set working directory for the build
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o /app/main ./cmd/cron

# Stage 2: Runtime
FROM alpine:latest

# Install Chromium dependencies dan timezone support
RUN apk update && apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont \
    libc6-compat \
    tzdata

# Set timezone ke Asia/Jakarta
ENV TZ=Asia/Jakarta

# (Opsional tapi disarankan) salin file timezone lokal
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# Set environment variable untuk Chromium
ENV CHROME_BIN=/usr/bin/chromium-browser
ENV DISPLAY=:0

# Set working directory
WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/main /main

# Copy template
COPY --from=builder /app/templates /app/templates

# Jalankan aplikasi
CMD ["/main"]
