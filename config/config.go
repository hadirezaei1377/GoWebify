package config

import (
	"log"
	"os"
)

var (
	ApiKey    string
	RedisAddr string
)

func LoadConfig() {
	ApiKey = os.Getenv("KAVEHNEGAR_API_KEY")
	RedisAddr = os.Getenv("REDIS_ADDR")
	if ApiKey == "" || RedisAddr == "" {
		log.Fatal("API key or Redis address not set")
	}
}
