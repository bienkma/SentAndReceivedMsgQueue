package config

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	HostName string
	Port     string
	QueueKafkaURL string
}

func AppConfig() Configuration {
	//file, ok := os.Open("/etc/app/app.json")
	file, ok := os.Open("app/config/app.json")
	defer file.Close()

	if ok != nil {
		panic("Can't open config file: app.json")
	}

	decoder := json.NewDecoder(file)
	cfg := Configuration{}

	if ok := decoder.Decode(&cfg); ok != nil {
		panic(ok)
	}
	return cfg
}
