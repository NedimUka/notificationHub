package services

import "log"

// InitMailgunService Initialize new mailgun service
func InitMailgunService(appConfig Config) {
	// TODO IMPLEMENT MAILGUN SERVICE
}

// SendEmailNotification -
func SendEmailNotification(subject, messagePalyload string, receivers []string) {

	for _, receiver := range receivers {
		// TODO Implement sending mobile notifications
		log.Printf("Send message to receiver %v with subject %v, and message %v", receiver, subject, messagePalyload)
	}
}
