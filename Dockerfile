FROM golang:latest

ENV USER root

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

RUN apt-get update && apt-get install -y \
    python3 \
 && rm -rf /var/lib/apt/lists/*

COPY . .

RUN go build
