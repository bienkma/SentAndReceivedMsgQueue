FROM golang:1.10.1

WORKDIR /go/src/app

RUN mkdir -p /etc/app

COPY app.json /etc/app/app.json

RUN apt-get install git

RUN \
    git clone https://github.com/edenhill/librdkafka.git /tmp/librdkafka && \
    cd /tmp/librdkafka && \
    ./configure && \
    make -j 8 && \
    make install && \
    ldconfig

RUN \
    git clone https://github.com/bienkma/SentAndRecivedMsgQueue.git /go/src/SentAndReceivedMsgQueue

EXPOSE 8080:8080
CMD ["app"]














