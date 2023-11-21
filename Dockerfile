# Golang Base Image
FROM golang:1.21.4-alpine3.18 AS build

## Build the executable in the first stage

WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Copy the source code files to the working directory
COPY src/ ./src

# Download the Go module dependencies
RUN go mod download

# Build the executable binary named "api" in the src directory
RUN go build -o api ./src

## Serve only the compiled binary in the second stage
FROM alpine:3.18.2 AS serve

## Necessary to run a health check in our docker-compose file
RUN apk --update --no-cache add curl

# Copy the pre-built binary file "api" from the build stage to the serve stage
COPY --from=build /app/api /app/api

# Set the command to run the "api" binary when the container starts
CMD ["/app/api"]