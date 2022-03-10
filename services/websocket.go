package services

import "log"

// InitWebsocketService Initialize new websocket service
func InitWebsocketService(appConfig Config) {
	// TODO IMPLEMENT FIREBASE SERVICE
}

// SendNotification -
func SendNotification(subject, messagePalyload string, receivers []string) {

	for _, receiver := range receivers {
		// TODO Implement sending mobile notifications
		log.Printf("Send message to receiver %v with subject %v, and message %v", receiver, subject, messagePalyload)
	}
}
