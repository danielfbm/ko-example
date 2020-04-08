FROM alpine

RUN apk add --no-cache docker curl git tar && \
    curl -OL https://github.com/openshift/source-to-image/releases/download/v1.2.0/source-to-image-v1.2.0-2a579ecd-linux-amd64.tar.gz && \
    tar -xvf source-to-image-v1.2.0-2a579ecd-linux-amd64.tar.gz && \
    rm source-to-image-v1.2.0-2a579ecd-linux-amd64.tar.gz && \
    chmod +x s2i && \
    cp s2i /usr/local/bin/s2i

ENTRYPOINT [ "/usr/local/bin/s2i" ]