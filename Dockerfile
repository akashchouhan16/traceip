# Use a base image with Go pre-installed
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the local machine to the container
COPY . .

# Build the Go application
RUN go build -o traceip .

# Set the command to run when the container starts
CMD ["./traceip"]
