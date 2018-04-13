package main

import (
	"fmt"
	"log"

	"net/http"
	"github.com/go-chi/chi"

	"github.com/bienkma/SentAndRecivedMsgQueue/app/config"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/router"
	"github.com/bienkma/SentAndRecivedMsgQueue/app/handler"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Read config file
	cfg := config.AppConfig()
	r := chi.NewRouter()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.QueueKafkaURL})
	if err != nil {
		panic(err)
	}


	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.QueueKafkaURL,
		"group.id":          "emtapto",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	p := handler.NewKafka{
		producer,
		consumer,
	}

	router.Register(r, p)
	log.Printf("Start listening on %s:%s", cfg.HostName, cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HostName, cfg.Port), r); err != nil {
		panic(err)
	}
}
