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

// add streaming
func PushToStream(msg model.MessageRequest, userID string) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: fmt.Sprintf("sms_stream_%s", userID), // stream based on user id
		Values: map[string]interface{}{
			"data": string(data),
		},
	}).Err()
	if err != nil {
		return err
	}
	fmt.Println("Message pushed to stream successfully")
	return nil
}

func ReadFromStream(userID string) (*model.MessageRequest, error) {
	streamName := fmt.Sprintf("sms_stream_%s", userID)

	result, err := RedisClient.XRead(ctx, &redis.XReadArgs{
		Streams: []string{streamName, "0"},
		Block:   0,
		Count:   1,
	}).Result()

	if err != nil || len(result) == 0 {
		return nil, err
	}

	var msg model.MessageRequest
	err = json.Unmarshal([]byte(result[0].Messages[0].Values["data"].(string)), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
