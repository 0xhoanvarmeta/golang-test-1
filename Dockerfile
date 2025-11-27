# Use minimal base image
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy pre-built binary from CI/CD or local build
# Binary should be built with: CGO_ENABLED=0 GOOS=linux go build -o bin/http-server ./cmd/http-server
COPY bin/http-server .

# Make binary executable
RUN chmod +x http-server

# Expose port (adjust if your app uses a different port)
EXPOSE 8080

# Run the binary
CMD ["./http-server"]
