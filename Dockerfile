# Use the official Golang image from the Docker Hub
FROM golang:1.12

WORKDIR /

# Copy all files from the current directory to the /app directory
COPY . .

# Build the application
RUN go build -o main .

# Run the application
CMD ["go run main"]