package handler

import (
	"github.com/bienkma/SentAndReceivedMsgQueue/app/view"
	"net/http"
)

func (kk *NewKafka) Reciver(ctx BaseHandler) view.ApiResponse {
	ctx.request.ParseForm()
	topic, _ := ctx.request.URL.Query()["topic"]

	if len(topic) != 0 {
		err := kk.NewConsumer.SubscribeTopics(topic, nil)
		defer kk.NewConsumer.Close()
		if err != nil {
			return view.ApiResponse{Status: http.StatusBadRequest, Msg: "Not msg"}
		}
		defer kk.NewConsumer.Close()
		msg := kk.NewConsumer.Poll(100)
		if err == nil {
			return view.ApiResponse{Status: http.StatusOK, Msg: string(msg.String())}
		}
	}
	return view.ApiResponse{Status: http.StatusBadRequest, Msg: "Not sent"}
}
