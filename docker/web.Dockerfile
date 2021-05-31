FROM golang:1.14

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

RUN apt-get install -y inotify-tools

RUN echo $GOPATH

COPY . /

ARG VERSION="4.13.0"

RUN set -x \
    && apt-get install -y git \
    && git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

WORKDIR /tmp/go-migrate

RUN set -x \
    && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
    && ./migrate -version

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /clean_web

RUN go mod download

RUN go get github.com/go-delve/delve/cmd/dlv

CMD sh /clean_web/docker/run.sh