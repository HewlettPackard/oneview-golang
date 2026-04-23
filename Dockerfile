FROM golang:1.20

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Proxy config (CRITICAL for your environment)
ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY

ENV HTTP_PROXY=$HTTP_PROXY
ENV HTTPS_PROXY=$HTTPS_PROXY
ENV http_proxy=$HTTP_PROXY
ENV https_proxy=$HTTPS_PROXY
ENV NO_PROXY=$NO_PROXY
ENV no_proxy=$NO_PROXY

# Install dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        gnupg \
        dirmngr \
        python3 \
        python3-pip && \
    ln -s /usr/bin/python3 /usr/bin/python && \
    pip3 install --no-cache-dir --upgrade pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN go build github.com/HewlettPackard/oneview-golang
