package handler

import (
	"net/http"
	"github.com/bienkma/SentAndReceivedMsgQueue/app/view"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type BaseHandler struct {
	response http.ResponseWriter
	request *http.Request
}

type NewKafka struct {
	NewProducer *kafka.Producer
	NewConsumer *kafka.Consumer
}

type HandleFunc func(ctx BaseHandler) view.ApiResponse

func GetBaseHandler(response http.ResponseWriter, request *http.Request) BaseHandler {
	return BaseHandler{response, request}
}


func MakeHandler(handlerFunc HandleFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		h := GetBaseHandler(response, request)
		res := handlerFunc(h)
		view.RenderJson(response, res)
	}
}