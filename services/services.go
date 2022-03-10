package services

// Config - app config
type Config struct {
	//TODO Implement app config
}

// Init will initialize services
func Init() {
	appConfig := Config{}
	// Initialize Firebase service
	InitFirebaseService(appConfig)
	// Initialize Mailgun Service
	InitMailgunService(appConfig)
	// Initialize SMS Service
	InitSMSService(appConfig)
	// Initialize Websocket
	InitWebsocketService(appConfig)
}
