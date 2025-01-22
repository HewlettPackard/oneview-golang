FROM golang:1.23

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . /go/src/github.com/HewlettPackard/oneview-golang
RUN go build github.com/HewlettPackard/oneview-golang
