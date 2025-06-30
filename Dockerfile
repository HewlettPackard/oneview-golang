FROM golang:1.21

ENV USER=root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Install Glide for dependency management
RUN go install github.com/Masterminds/glide@latest

COPY . /go/src/github.com/HewlettPackard/oneview-golang

# Install dependencies using Glide
RUN glide install

# Since this is a library, we'll just run tests to verify the build
RUN go test ./...
