package notify

import (
	"fmt"
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
	// slack
	slackWebhook := os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK")
	if len(slackWebhook) > 0 {
		sendSlackWebhook(
			fmt.Sprintf(
				"[Info] %v %v",
				getNotifySourceLabelString(),
				message,
			),
			os.Getenv("NOTIFY_INFO_SLACK_WEBHOOK"),
		)
	}

	// discord
	discordWebhook := os.Getenv("NOTIFY_INFO_DISCORD_WEBHOOK")
	if len(discordWebhook) > 0 {
		_ = sendDiscordWebhook(
			fmt.Sprintf(
				"[Info] %v %v",
				getNotifySourceLabelString(),
				message,
			),
			os.Getenv("NOTIFY_INFO_DISCORD_WEBHOOK"),
		)
	}
}

// Alert sends a notification to an alert channel
func Alert(message string) {
	slackWebhook := os.Getenv("NOTIFY_ALERT_SLACK_WEBHOOK")
	if len(slackWebhook) > 0 {
		sendSlackWebhook(
			fmt.Sprintf(
				"[ALERT] %v %v",
				getNotifySourceLabelString(),
				message,
			),
			os.Getenv("NOTIFY_ALERT_SLACK_WEBHOOK"),
		)
	}

	// discord
	discordWebhook := os.Getenv("NOTIFY_ALERT_DISCORD_WEBHOOK")
	if len(discordWebhook) > 0 {
		_ = sendDiscordWebhook(
			fmt.Sprintf(
				"[ALERT] %v %v",
				getNotifySourceLabelString(),
				message,
			),
			os.Getenv("NOTIFY_ALERT_DISCORD_WEBHOOK"),
		)
	}
}
