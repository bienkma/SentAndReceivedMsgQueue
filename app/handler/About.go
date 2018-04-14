package handler

import (
	"net/http"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/view"
)

// Intro About me. So it's not necessary BaseHandler
func About(_ BaseHandler) view.ApiResponse{
	return view.ApiResponse{Status:http.StatusOK, Msg:"Welcome simple chat system!"}
}