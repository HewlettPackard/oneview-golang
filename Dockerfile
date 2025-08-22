FROM golang:1.11

# Install Python3 + pip (needed for scripts/tests/tools)
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    rm -rf /var/lib/apt/lists/*

ENV USER root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . .

RUN go build github.com/HewlettPackard/oneview-golang

CMD ["/bin/bash"]
