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
  export GOPATH=`$home/go` && \
  mkdir /etc/app/

COPY app.json /etc/app/app.json

RUN cd /tmp/ && \
    git clone https://github.com/bienkma/SentAndRecivedMsgQueue.git /tmp/SentAndReceivedMsgQueue && \
    cd /tmp/SentAndReceivedMsgQueue














