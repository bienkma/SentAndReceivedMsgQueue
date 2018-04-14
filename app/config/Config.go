package config

import "os"

type Configuration struct {
	HostName string
	Port     string
	QueueKafkaURL string
}

func AppConfig() Configuration {
	return Configuration{
		HostName: os.Getenv("APP_BIND_IP"),
		Port: os.Getenv("8080"),
		QueueKafkaURL: os.Getenv("BROKER_LIST_URL"),
	}
}
