FROM golang:1.11

ENV USER root
ARG http_proxy
ARG https_proxy
ARG no_proxy

ENV HTTP_PROXY=${http_proxy}
ENV HTTPS_PROXY=${https_proxy}
ENV NO_PROXY=${no_proxy}

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# RUN apt-get update && apt-get install -y \
#     python3 \
#  && rm -rf /var/lib/apt/lists/*
RUN apt update -y

RUN apt install python3-pip -y



COPY . /go/src/github.com/HewlettPackard/oneview-golang
RUN go build github.com/HewlettPackard/oneview-golang

