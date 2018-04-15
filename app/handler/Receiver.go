package handler

import (
	"github.com/bienkma/SentAndReceivedMsgQueue/app/view"
	"net/http"
	"time"
)

func (kk *NewKafka) Reciver(ctx BaseHandler) view.ApiResponse {
	ctx.request.ParseForm()
	topic, _ := ctx.request.URL.Query()["topic"]
	if len(topic) != 0 {
		kk.NewConsumer.SubscribeTopics([]string{"roomchat"}, nil)
		msg, err := kk.NewConsumer.ReadMessage(100 * time.Millisecond)
		if err == nil{
			return view.ApiResponse{Status: http.StatusOK, Msg: string(msg.Value)}
		}
	}
	return view.ApiResponse{Status: http.StatusBadRequest, Msg: "Not sent"}
}
