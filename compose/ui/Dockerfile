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

#ENV TOPIC=roomchat
#ENV API_SENDER_URL=http://192.168.1.74:8080/apis/sender
#ENV API_RECEIVER_URL=http://192.168.1.74:8080/apis/receiver?topic=roomchat

CMD ["ui"]
