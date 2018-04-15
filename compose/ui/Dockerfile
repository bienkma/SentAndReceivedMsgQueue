FROM golang:1.10.1

WORKDIR /go/src/app

RUN apt-get install git

RUN \
    git clone https://github.com/bienkma/SentAndRecivedMsgQueue.git /go/src/app/SentAndReceivedMsgQueue && \
    cd /go/src/app/SentAndReceivedMsgQueue/ui && \
    go get -v -d . && \
    go install && \
    cp /go/bin/ui /go/src/app/ui && \
    rm -rf /go/src/app/SentAndReceivedMsgQueue

CMD ["ui"]
