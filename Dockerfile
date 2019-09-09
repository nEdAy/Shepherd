# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="nEdAy <shengsu15@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o Shepherd .

# Expose port 8443 to the outside world
EXPOSE 8443

# Command to run the executable
CMD ["Shepherd"]