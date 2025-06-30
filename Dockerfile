FROM golang:1.21

ENV USER=root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . /go/src/github.com/HewlettPackard/oneview-golang
RUN go build .
