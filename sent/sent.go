package main

import (
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func sentMsg(sender *kafka.Producer,brokerLists string, topic string, msg []byte) error{
	err := sender.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: msg,
	}, nil)
	if err != nil{
		return err
	}

	// Wait for deliveries msg 5 second
	sender.Flush(5 * 1000)
	return nil
}

func main()  {
	var brokerLists string = "127.0.0.1:6667"
	var topic string = "chatQueue"

	sender, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerLists})
	if err != nil {
		log.Printf("Can't connection to kafka: %s", err)
	}
	if err := sentMsg(sender, brokerLists, topic, []byte("one")); err != nil{
		log.Println("Can't sent message to topic %s with error %s", topic, err)
	}
}