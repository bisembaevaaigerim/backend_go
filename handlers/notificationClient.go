package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

var restyClient = resty.New()

func init() {
	restyClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Printf("[Resty] --> %s %s", req.Method, req.URL)
		return nil
	})

	restyClient.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Printf("[Resty] <-- %d from %s", resp.StatusCode(), resp.Request.URL)
		return nil
	})
}

func SendBookingNotification(userID uint, eventID uint) error {
	notifyURL := getNotifyURL()
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"user_id":  userID,
			"event_id": eventID,
			"message":  fmt.Sprintf("Your booking for event #%d is confirmed!", eventID),
		}).
		Post(notifyURL + "/notify")

	if err != nil {
		return fmt.Errorf("notification request failed: %w", err)
	}

	log.Printf("[Resty] Notification response: %s", resp.String())
	return nil
}
func getNotifyURL() string {
	if url := os.Getenv("NOTIFICATION_SERVICE_URL"); url != "" {
		return url
	}
	return "http://localhost:8085"
}
