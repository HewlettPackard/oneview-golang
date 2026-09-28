FROM golang:1.11

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# SDK Automator CI runs Python inside this image, so python3 + pip3 are required
RUN apt-get update && apt-get install -y --no-install-recommends python3 python3-pip \
    && rm -rf /var/lib/apt/lists/*

COPY . /go/src/github.com/HewlettPackard/oneview-golang
RUN go build github.com/HewlettPackard/oneview-golang
