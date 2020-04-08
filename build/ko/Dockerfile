FROM golang:1.13-alpine

RUN apk add --no-cache docker curl git

RUN GO111MODULE=on go get github.com/google/ko/cmd/ko
RUN GO111MODULE=on go get -u github.com/google/go-containerregistry/cmd/crane

ENTRYPOINT ["/go/bin/ko"]
