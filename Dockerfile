# Start with Go base image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy application files
COPY . .

# Build the Go application
RUN go build -o web-app .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./web-app"]