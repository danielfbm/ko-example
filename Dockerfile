FROM index.alauda.cn/library/golang:1.13-alpine as builder

ENV GO111MODULE=on
WORKDIR /workspace
COPY . /workspace

RUN ls -lha && CGO_ENABLED=0 go build -o bin/app .

FROM index.alauda.cn/library/alpine:3.11

COPY --from=builder /workspace/bin/app /app

ENTRYPOINT ["/app"]
