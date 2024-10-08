package model

// model for requesting sms
type MessageRequest struct {
	PhoneNumbers []string `json:"phone_numbers"`
	Message      string   `json:"message"`
}
