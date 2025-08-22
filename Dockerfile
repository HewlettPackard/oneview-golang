FROM golang:1.21-bullseye

RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    rm -rf /var/lib/apt/lists/*

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . /go/src/github.com/HewlettPackard/oneview-golang

RUN go build github.com/HewlettPackard/oneview-golang

CMD ["/bin/bash"]

