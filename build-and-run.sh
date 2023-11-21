#!/bin/bash

# Build the Docker image
docker build -t kts-backend .

# Run the Docker container
docker run -p 8080:8080 kts-backend