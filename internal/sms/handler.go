package sms

import (
	"GoWebify/internal/model"
	"encoding/json"
	"net/http"
)

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.PhoneNumbers) == 0 || req.Message == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = PushToQueue(req)
	if err != nil {
		http.Error(w, "Failed to queue SMS", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SMS queued successfully"))
}
