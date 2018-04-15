FROM golang:1.10.1

WORKDIR /go/src/app

RUN apt-get install git

RUN \
    git clone https://github.com/edenhill/librdkafka.git /tmp/librdkafka && \
    cd /tmp/librdkafka && \
    git checkout tags/v0.11.4 && \
    ./configure && \
    make -j 8 && \
    make install && \
    ldconfig

RUN \
    git clone https://github.com/bienkma/SentAndRecivedMsgQueue.git /go/src/app/SentAndReceivedMsgQueue && \
    cd /go/src/app/SentAndReceivedMsgQueue/app && \
    go get -v -d . && \
    go install
COPY start.sh /bin/start.sh
RUN chmod 755 /bin/start.sh

CMD ["/bin/start.sh"]
