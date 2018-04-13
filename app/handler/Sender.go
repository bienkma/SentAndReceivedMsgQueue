package handler

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/view"
	"net/http"
)

func (kk *NewKafka) Sender(ctx BaseHandler) view.ApiResponse {
	ctx.request.ParseForm()
	topic := ctx.request.Form.Get("topic")
	msg := ctx.request.Form.Get("msg")
	if topic != "" && msg != "" {
		err := kk.NewProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			return view.ApiResponse{Code: http.StatusBadRequest, Data: "Not sent"}
		}
		return view.ApiResponse{Code: http.StatusOK, Data: "sent: " + msg + " to channel: " + topic}
	}
	return view.ApiResponse{Code: http.StatusBadRequest, Data: "Not sent"}
}
