package handler

import (
	"net/http"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/view"
)

// Intro About me. So it's not necessary BaseHandler
func About(_ BaseHandler) view.ApiResponse{
	return view.ApiResponse{Code:http.StatusOK, Data:"Welcome simple chat system!"}
}