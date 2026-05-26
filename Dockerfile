# PQC Story 1: Go 1.26+ provides finalized FIPS 203 PQC hybrid key exchange (X25519MLKEM768)
# in crypto/tls, enabled by default. No OpenSSL dependency needed.
# See: PQC Enablement Checklist Section 2.1.4
FROM golang:1.26

WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install Python3 + pip for SDK Automator CI
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3 python3-pip && \
    rm -rf /var/lib/apt/lists/*

COPY . .
RUN go mod vendor && go build ./...
