FROM golang:1.11

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install Python3 and pip
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    ln -s /usr/bin/python3 /usr/bin/python && \
    pip3 install --upgrade pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy repo
COPY . /go/src/github.com/HewlettPackard/oneview-golang

# Build Go SDK
RUN go build github.com/HewlettPackard/oneview-golang
