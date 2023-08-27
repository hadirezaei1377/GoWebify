package main

import (
	"GoWebify/internal/models"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// function for listennig to channel
func ListenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			SendMsg(msg)
		}
	}()
}

func SendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()

	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, "Hello,<strong>world</strong>!")

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent!")
	}
}
