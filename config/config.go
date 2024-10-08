package config

import (
	"log"
	"os"
)

var (
	ApiKey string
)

func LoadConfig() {
	ApiKey = os.Getenv("KAVEHNEGAR_API_KEY")
	if ApiKey == "" {
		log.Fatal("API key not set")
	}
}
