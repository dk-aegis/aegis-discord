# Start from an official Golang base image
FROM golang:1.22.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project to the working directory inside the container
COPY . .

# Build the Go binary
RUN go build -o aegis-discord .

# Run the compiled Go binary
CMD ["/app/aegis-discord"]
