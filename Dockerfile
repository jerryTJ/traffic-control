# Stage 1: Build Stage --platform linux/arm64
FROM golang:1.23.3-alpine3.20 as builder
# Set the working directory
WORKDIR /app
# Copy the application source code
COPY . .
# Copy dependencies and download modules
COPY go.mod go.sum ./
RUN go mod download
# Build the application
RUN go build -o controller Application.go
# Stage 2: Runtime Stage
FROM  alpine:3.20.3
# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates
# Set the working directory
WORKDIR /app/
# Copy the compiled binary from the build stage
# COPY --from=builder /app/deployments/tini /app/
COPY --from=builder /app/controller /app/controller
# Command to run the application
# ENTRYPOINT [ "/app/tini", "--" ]
CMD ["/app/controller"]
