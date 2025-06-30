FROM golang:1.20

# Pass proxy from environment to Docker build
ARG http_proxy
ARG https_proxy
ENV http_proxy=${http_proxy}
ENV https_proxy=${https_proxy}

# Set working directory
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Use the proxy for apt too
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    apt-get clean

# Copy project
COPY . .

# Build the Go app
RUN go build -o oneview-golang .

CMD ["./oneview-golang"]
