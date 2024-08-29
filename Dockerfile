FROM golang:1.20

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . .
RUN ls -la /go/src/github.com/HewlettPackard/oneview-golang
RUN go build github.com/HewlettPackard/oneview-golang
