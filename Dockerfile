FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download || true  # `|| true` to avoid failure if go.sum is missing

COPY . .

RUN go build -o oneview-golang

CMD ["./oneview-golang"]
