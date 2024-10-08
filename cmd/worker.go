package main

import (
	"GoWebify/internal/sms"
	"log"
	"time"
)

func queue() {
	sms.InitQueue()

	log.Println("Worker is running and waiting for messages...")

	for {

		userID := "example_user_id"
		msg, err := sms.ReadFromStream(userID)
		if err != nil {
			log.Println("No messages in stream for user", userID, "waiting...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = sms.SendSms(*msg)
		if err != nil {
			log.Println("Failed to send SMS for user", userID, ":", err)
		} else {
			log.Println("SMS sent successfully for user", userID)
		}
	}
}
