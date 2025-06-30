FROM golang:1.11
 
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang
COPY . .
 
RUN go mod tidy
RUN go build -o oneview-golang