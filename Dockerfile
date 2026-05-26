# PQC Story 1+3: Go 1.24+ provides native PQC hybrid key exchange in crypto/tls.
# No OpenSSL dependency needed.
# See: PQC Enablement Checklist Section 1.1.1, 2.1.4
FROM golang:1.24

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install Python3 + pip for SDK Automator CI
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3 python3-pip && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN go mod vendor && go build ./...
