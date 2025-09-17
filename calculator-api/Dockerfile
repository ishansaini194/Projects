# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o calculator-api ./cmd/server

# Final stage: small runtime image
FROM alpine:3.20

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/calculator-api .

# Run on container start
EXPOSE 8080
CMD ["./calculator-api"]
