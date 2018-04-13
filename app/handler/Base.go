package handler

import (
	"net/http"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/view"
)

type BaseHandler struct {
	w http.ResponseWriter
	r *http.Request
}

type HandleFunc func(ctx BaseHandler) view.ApiResponse

func GetBaseHandler(w http.ResponseWriter, r *http.Request) BaseHandler {
	return BaseHandler{w, r}
}


func MakeHandler(handlerFunc HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := GetBaseHandler(w, r)
		res := handlerFunc(h)
		view.RenderJson(w, res)
	}
}