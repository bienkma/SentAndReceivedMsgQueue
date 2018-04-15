package handler

import (
	"github.com/bienkma/SentAndReceivedMsgQueue/app/view"
	"net/http"
)

func (kk *NewKafka) Reciver(ctx BaseHandler) view.ApiResponse {
	ctx.request.ParseForm()
	topic, _ := ctx.request.URL.Query()["topic"]
	if len(topic) != 0 {
		if err := kk.NewConsumer.SubscribeTopics([]string{topic[0], "^aRegex.*[Tt]opic"}, nil); err != nil {
			return view.ApiResponse{Status: http.StatusBadRequest, Msg: "Not msg"}
		}
		defer kk.NewConsumer.Close()
		msg, err := kk.NewConsumer.ReadMessage(1)
		if err == nil {
			return view.ApiResponse{Status: http.StatusOK, Msg: string(msg.Value)}
		}
	}
	return view.ApiResponse{Status: http.StatusBadRequest, Msg: "Not sent"}
}
