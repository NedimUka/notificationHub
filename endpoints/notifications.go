package endpoints

import (
	"context"
	"log"

	pb "github.com/NedimUka/notification-hub/protoModels"
	services "github.com/NedimUka/notification-hub/services"
)

// SendNotification is used for sending any app notifications, emails or sms, to users
func (*Server) SendNotification(ctx context.Context, sendNotificationRequest *pb.NotificationRequest) (*pb.NotificationResponse, error) {

	// Check notification type and run correct service
	switch sendNotificationRequest.NotificationType {
	case pb.NotificationRequest_EMAIL_NOTIFICATION:
		log.Printf("Sending Email notification")
		services.SendEmailNotification(sendNotificationRequest.Subject, sendNotificationRequest.MessagePayload, sendNotificationRequest.Receivers)
	case pb.NotificationRequest_MOBILE_NOTIFICATION:
		log.Printf("Sending Email notification")
		services.SendMobileNotification(sendNotificationRequest.Subject, sendNotificationRequest.MessagePayload, sendNotificationRequest.Receivers)
	case pb.NotificationRequest_SMS_NOTIFICATION:
		log.Printf("Sending Email notification")
		services.SendSMSNotification(sendNotificationRequest.Subject, sendNotificationRequest.MessagePayload, sendNotificationRequest.Receivers)
	case pb.NotificationRequest_WEBAPP_NOTIFICATION:
		log.Printf("Sending Email notification")
		services.SendNotification(sendNotificationRequest.Subject, sendNotificationRequest.MessagePayload, sendNotificationRequest.Receivers)

	}

	return &pb.NotificationResponse{}, nil

}
