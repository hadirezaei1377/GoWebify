package sms

import (
	"context"
	"fmt"
	"log"

	"GoWebify/config"
	"GoWebify/internal/model"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// store redis connection
var RedisClient *redis.Client

// start connecting to redis
func InitQueue() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

// push sms to queue
func PushToQueue(msg model.MessageRequest) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := RedisClient.LPush(ctx, "sms_queue", data).Err(); err != nil {
		return err
	}
	fmt.Println("Message pushed to queue successfully")
	return nil
}

func PopFromQueue() (*model.MessageRequest, error) {
	data, err := RedisClient.RPop(ctx, "sms_queue").Result()
	if err != nil {
		return nil, err
	}

	var msg model.MessageRequest
	err = json.Unmarshal([]byte(data), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
