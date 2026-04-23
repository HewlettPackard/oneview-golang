FROM golang:1.20

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    ln -s /usr/bin/python3 /usr/bin/python

COPY . .

RUN go build github.com/HewlettPackard/oneview-golang
