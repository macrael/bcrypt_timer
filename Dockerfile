# This container runs a go program timing how long the bcrypt library takes
# for given work factors.

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10.0 AS build

RUN go get github.com/golang/dep/cmd/dep

COPY ./ /go/src/github.com/macrael/bcrypt_timer/
WORKDIR /go/src/github.com/macrael/bcrypt_timer/
RUN dep ensure -vendor-only

# Install tools required to build the project
# These linker flags create a standalone binary that will run in scratch.
RUN go build -o /bin/bcrypt_timer -ldflags "-linkmode external -extldflags -static" .

# We don't need any of the Go mechanics, just the binaries. Copy into a single
# layer image.
FROM gcr.io/distroless/base
COPY --from=build /bin/bcrypt_timer /bin/bcrypt_timer

ENTRYPOINT ["/bin/bcrypt_timer"]
