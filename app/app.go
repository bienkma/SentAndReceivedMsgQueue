package main

import (
	"fmt"
	"log"

	"net/http"
	"github.com/go-chi/chi"

	"github.com/bienkma/SentAndRecivedMsgQueue/app/config"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/router"
)

func main() {
	// Read config file
	cfg := config.AppConfig()
	r := chi.NewRouter()

	router.Register(r)
	log.Printf("Start listening on %s:%s", cfg.HostName, cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HostName, cfg.Port), r); err != nil {
		panic(err)
	}
}
