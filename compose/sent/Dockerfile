FROM ubuntu:16.04
MAINTAINER bienkma (bientd88@gmail.com)

RUN mkdir -p $home/go/src

RUN \
  apt-get update && \
  apt-get install -y --no-install-recommends build-essential libfuse-dev libcurl4-openssl-dev automake libtool pkg-config libssl-dev wget tar git && \
  apt-get clean && \
  apt-get autoremove -y && \
  rm -rf /var/lib/{apt,dpkg,cache,log}/

RUN \
  export GOROOT='/usr/local/go' && \
  export GOPATH=`$home/go`

COPY sent/ sent

RUN \
    cd sent && \
    go get all
    go build -o sent sent.go
