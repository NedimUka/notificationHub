package services

import "log"

// InitFirebaseService Initialize new firebase service
func InitFirebaseService(appConfig Config) {
	// TODO IMPLEMENT FIREBASE SERVICE

}

// SendMobileNotification -
func SendMobileNotification(subject, messagePalyload string, receivers []string) {

	for _, receiver := range receivers {
		// TODO Implement sending mobile notifications
		log.Printf("Send message to receiver %v with subject %v, and message %v", receiver, subject, messagePalyload)
	}
}
