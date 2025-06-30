FROM python:3.11-slim

# Install Go and Python build tools
RUN apt-get update && \
    apt-get install -y wget git tar build-essential python3 python3-pip && \
    wget https://golang.org/dl/go1.11.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.11.linux-amd64.tar.gz && \
    rm go1.11.linux-amd64.tar.gz && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set Go environment
ENV PATH="/usr/local/go/bin:$PATH"
ENV GOPATH=/go
ENV PATH="$GOPATH/bin:$PATH"

RUN mkdir -p $GOPATH

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . .

RUN go build .
