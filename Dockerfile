FROM golang:1.11

ENV USER root
ARG http_proxy
ARG https_proxy
ARG no_proxy

ENV HTTP_PROXY=${http_proxy}
ENV HTTPS_PROXY=${https_proxy}
ENV NO_PROXY=${no_proxy}

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install Python 3 and pip
RUN apt-get update -y && \
    apt-get install -y python3 python3-pip && \
    ln -s /usr/bin/python3 /usr/bin/python && \
    ln -s /usr/bin/pip3 /usr/bin/pip && \
    rm -rf /var/lib/apt/lists/*

COPY . /go/src/github.com/HewlettPackard/oneview-golang

# Build Go project (optional)
RUN go build github.com/HewlettPackard/oneview-golang

