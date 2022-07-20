package notify

import (
	"fmt"
	"log"
	"os"
)

// getNotifySourceLabel provides a description label for where the notification
// is being sent from.
// for example of you are sending backups from host-1 the environment variable
// would override the default (hostname)
func getNotifySourceLabel() string {
	notifySourceLabel := os.Getenv("NOTIFY_SOURCE_LABEL")
	if len(notifySourceLabel) > 0 {
		return notifySourceLabel
	}

	hostname, _ := os.Hostname()
	if len(hostname) > 0 {
		return hostname
	}

	return ""
}

// getNotifySourceLabelString provides a formatted source label for messages
func getNotifySourceLabelString() string {
	label := getNotifySourceLabel()
	if len(label) > 0 {
		return fmt.Sprintf("[%v]", label)
	}

	return ""
}

// Info sends a notification to an info channel
func Info(message string) {
	slackWebhook := os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK")
	if len(slackWebhook) > 0 {
		err := SendSlackNotification(&SendSlackNotificationInput{
			WebhookURL: os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK"),
			Message: fmt.Sprintf(
				"[Info] %v %v",
				getNotifySourceLabelString(),
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
				"[ALERT] %v %v",
				getNotifySourceLabelString(),
				message,
			),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
