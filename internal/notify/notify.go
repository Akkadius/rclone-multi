package notify

import (
	"fmt"
	"log"
	"os"
)

// Info sends a notification to an info channel
func Info(message string) {
	slackWebhook := os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK")
	if len(slackWebhook) > 0 {
		err := SendSlackNotification(&SendSlackNotificationInput{
			WebhookURL: os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK"),
			Message: fmt.Sprintf(
				"[%v] %v",
				os.Getenv("BACKUP_DEPLOYMENT_NAME"),
				message,
			),
		})
		if err != nil {
			log.Println(err)
		}
	}
}

// Alert sends a notification to an alert channel
func Alert(message string) {
	slackWebhook := os.Getenv("NOTIFY_ALERT_SLACK_WEBHOOK")
	if len(slackWebhook) > 0 {
		err := SendSlackNotification(&SendSlackNotificationInput{
			WebhookURL: os.Getenv("NOTIFY_ALERT_SLACK_WEBHOOK"),
			Message: fmt.Sprintf(
				"[%v] %v",
				os.Getenv("BACKUP_DEPLOYMENT_NAME"),
				message,
			),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
