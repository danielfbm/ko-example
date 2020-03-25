FROM golang:1.13-alpine

RUN apk add --no-cache docker curl

RUN GO111MODULE=on go get github.com/google/ko/cmd/ko
RUN GO111MODULE=on go get -u github.com/google/go-containerregistry/cmd/crane


