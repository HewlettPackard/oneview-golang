# PQC Story 1+3: Chainguard Go image with Go 1.24+ (native PQC hybrid key exchange)
# Go's crypto/tls provides PQC support natively — no OpenSSL dependency needed.
# See: PQC Enablement Checklist Section 1.1.1, 2.1.4
# Go modules enabled — no GOPATH layout required.
FROM cgr.dev/chainguard/go:latest

WORKDIR /app

COPY . .
RUN go mod vendor && go build ./...
