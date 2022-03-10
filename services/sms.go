package services

import "log"

// InitSMSService Initialize new SMS service
func InitSMSService(appConfig Config) {
	// TODO IMPLEMENT SMS SERVICE
}

// SendSMSNotification -
func SendSMSNotification(subject, messagePalyload string, receivers []string) {

	for _, receiver := range receivers {
		// TODO Implement sending mobile notifications
		log.Printf("Send message to receiver %v with subject %v, and message %v", receiver, subject, messagePalyload)
	}
}
