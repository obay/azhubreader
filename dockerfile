# Start from the official Golang image
FROM golang:1.22.5-alpine3.20

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o azhubreader .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./azhubreader"]