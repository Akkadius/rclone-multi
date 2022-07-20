package notify

import (
	"fmt"
	"log"
	"os"
)

func Send(message string) {
	slackWebhook := os.Getenv("BACKUP_SLACK_NOTIFY_WEBHOOK")
	if len(slackWebhook) > 0 {
		err := SendSlackNotification(&SendSlackNotificationInput{
			WebhookURL: os.Getenv("BACKUP_SLACK_NOTIFY_WEBHOOK"),
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
