package main

import (
	"GoWebify/internal/sms"
	"log"
	"time"
)

func queued() {
	sms.InitQueue()

	log.Println("Worker is running and waiting for messages...")

	for {

		msg, err := sms.PopFromQueue()
		if err != nil {
			log.Println("No messages in queue, waiting...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = sms.SendSms(*msg)
		if err != nil {
			log.Println("Failed to send SMS:", err)
		} else {
			log.Println("SMS sent successfully")
		}
	}
}
