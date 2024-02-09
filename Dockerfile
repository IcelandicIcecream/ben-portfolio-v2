# Use golang:latest as the build env for the build stage
FROM golang:latest AS build-stage

# Set working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files, migrations files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . /app

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

# Use alpine:latest for the final image
FROM alpine:latest

# Add bash for convenience
RUN apk add --no-cache bash

# Set working directory in the final image
WORKDIR /app

# Copy the built binary and configuration file from the build environment
COPY --from=build-stage /app/main .
COPY --from=build-stage /app/views ./views
COPY --from=build-stage /app/styles ./styles
COPY --from=build-stage /app/assets ./assets
COPY --from=build-stage /app/.env .

# Expose port 8080 (if your app uses this port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
