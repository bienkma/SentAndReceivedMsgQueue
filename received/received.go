package main

import (
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func receivedMsg(receiver *kafka.Consumer, brokerLists string, topic string, group string) error {
	err := receiver.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := receiver.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			return err
		}
	}
	return nil
}