FROM golang:1.21

ENV USER=root
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

COPY . /go/src/github.com/HewlettPackard/oneview-golang

# Since this is a library project with vendored dependencies, run tests to verify the build
# Use -mod=vendor to use the vendored dependencies
RUN go test -mod=vendor ./...
