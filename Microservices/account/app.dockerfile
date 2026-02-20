# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install certs (if your app makes HTTPS calls)
RUN apk add --no-cache ca-certificates

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./account/cmd/account

# ---- Final Stage ----
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]