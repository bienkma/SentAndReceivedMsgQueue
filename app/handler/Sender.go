package handler

import (
	"net/http"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/view"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/config"
)

func store2Kafka(p *kafka.Producer, topic string, msg []byte) error{
	return p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition:kafka.PartitionAny},
		Value: msg,
	}, nil)
}

func Sender(ctx BaseHandler) view.ApiResponse {
	ctx.r.ParseForm()
	msg := ctx.r.FormValue("msg")
	topic := ctx.r.FormValue("topic")
	p := kafka.Producer{&kafka.ConfigMap{"bootstrap.servers": config.AppConfig().QueueKafkaURL}}
	if err:= store2Kafka(p, topic, []byte(msg)); err != nil{
		return view.ApiResponse{Code: http.StatusOK, Data: "sent"}
	}
	return view.ApiResponse{Code: http.StatusBadRequest, Data: "Not sent"}
}
