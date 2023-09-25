# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and Go sum files to the container
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# You can add a script (wait-for-postgres.sh) to wait for PostgreSQL to start
# This script is useful to ensure the database is available before starting the API
# Here's a simple example of wait-for-postgres.sh:
# #!/bin/sh
# set -e
# host="$1"
# shift
# until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
#   >&2 echo "PostgreSQL is unavailable - sleeping"
#   sleep 1
# done
# >&2 echo "PostgreSQL is up - executing command"
