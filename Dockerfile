FROM golang:1.11

RUN apt-get update && apt-get install -y --no-install-recommends \
      python3 python3-pip git jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . .

RUN go build github.com/HewlettPackard/oneview-golang

CMD ["/bin/bash"]
