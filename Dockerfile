FROM golang:latest

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application

RUN go mod tidy

RUN go mod vendor

RUN go build -o oneview-golang .

# Optionally set the entrypoint
CMD ["./oneview-golang"]
