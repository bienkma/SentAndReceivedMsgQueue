package router

import (
	"github.com/sirupsen/logrus"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"

	myMiddleware "github.com/bienkma/SentAndRecivedMsgQueue/app/router/middleware"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/handler"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/log"
)

func Register(r *chi.Mux) {

	// Setup log
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	// Add Middleware for router
	r.Use(chiMiddleware.Compress(2, "gzip"))
	r.Use(myMiddleware.CORS)

	r.Use(log.NewStructuredLogger(logger))

	// api router
	r.Group(func(r chi.Router) {
		r.Get("/apis/about", handler.MakeHandler(handler.About))
		r.Post("/apis/sender", handler.MakeHandler(handler.Sender))
		r.Get("/apis/receiver", handler.MakeHandler(handler.receiver))
	})
}