package config

import "os"

type Configuration struct {
	Port     string
	QueueKafkaURL string
}

func AppConfig() Configuration {
	return Configuration{
		Port: os.Getenv("API_BIND_PORT"),
		QueueKafkaURL: os.Getenv("BROKER_LIST_URL"),
	}
}
