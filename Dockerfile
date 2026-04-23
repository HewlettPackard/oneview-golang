FROM golang:1.20

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Fix apt + install required packages (handles GPG + cert issues)
RUN apt-get update --allow-releaseinfo-change && \
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

# Copy repo
COPY . .

# Build Go SDK
RUN go build github.com/HewlettPackard/oneview-golang
