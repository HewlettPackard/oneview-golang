FROM golang:1.11

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install python3 and pip
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY . /go/src/github.com/HewlettPackard/oneview-golang
RUN go build github.com/HewlettPackard/oneview-golang
