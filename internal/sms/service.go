package sms

import (
	"GoWebify/config"
	"GoWebify/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type kavehNegarResponse struct {
	Status int `json:"return"`
}

func SendSms(req model.MessageRequest) error {
	apiUrl := "https://api.kavenegar.com/v1/" + config.ApiKey + "/sms/send.json"

	// request body
	payload, _ := json.Marshal(map[string]interface{}{
		"receptor": req.PhoneNumbers,
		"message":  req.Message,
	})

	httpClient := &http.Client{Timeout: 10 * time.Second}

	// retry pattern for ensuring of sending
	return retry(3, 2*time.Second, func() error {
		resp, err := httpClient.Post(apiUrl, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to send SMS: status code %d", resp.StatusCode)
		}

		var kavehResp kavehNegarResponse
		if err := json.NewDecoder(resp.Body).Decode(&kavehResp); err != nil {
			return err
		}

		if kavehResp.Status != 200 {
			return fmt.Errorf("failed to send SMS: KavehNegar status %d", kavehResp.Status)
		}

		return nil
	})
}
