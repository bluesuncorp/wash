package globals

import "github.com/go-playground/log"

//EmailSettings contains all email settings related information
type EmailSettings struct {
	SMTPServer   string
	SMTPUsername string
	SMTPPassword string
	ContactEmail string
	SMTPPort     int
}

// NewEmail returns a new email instance
func NewEmail(smtpServer string, smtpUsername string, smtpPassword string, smtpPort int, contactEmail string) *EmailSettings {

	log.Info("Initializing Email Settings")

	return &EmailSettings{
		SMTPServer:   smtpServer,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
		ContactEmail: contactEmail,
		SMTPPort:     smtpPort,
	}
}
