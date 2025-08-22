FROM golang:1.22 AS builder

# Install Python3 + pip (if still needed for tests/tools)
RUN apt-get update && apt-get install -y python3 python3-pip && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o oneview-golang .

# Final lightweight runtime image
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y python3 python3-pip && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/oneview-golang .

CMD ["./oneview-golang"]
