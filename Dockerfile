FROM golang:1.20

# Set environment variables
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Copy go.mod and go.sum separately for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o oneview-golang

# Optional: set default command
CMD ["./oneview-golang"]
