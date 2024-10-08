package main

import (
	"GoWebify/config"
	"GoWebify/internal/sms"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/send-sms", sms.SmsHandler)

	log.Println("SMS Service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
