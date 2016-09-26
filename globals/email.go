package globals

import (
	"github.com/go-playground/log"
)

// EmailSettings contains all email settings related information
type EmailSettings interface {
	SMTPServer() string
	SMTPUsername() string
	SMTPPassword() string
	ContactEmail() string
	SMTPPort() int
}

// emailSettings contains all email settings related information
type emailSettings struct {
	smtpServer   string
	smtpUsername string
	smtpPassword string
	contactEmail string
	smtpPort     int
}

var _ EmailSettings = new(emailSettings)

// newEmail returns a new email instance
func newEmail(smtpServer string, smtpUsername string, smtpPassword string, smtpPort int, contactEmail string) EmailSettings {

	log.Info("Initializing Email Settings")

	return &emailSettings{
		smtpServer:   smtpServer,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
		contactEmail: contactEmail,
		smtpPort:     smtpPort,
	}
}

// SMTPServer returns the applications SMTP server
func (e *emailSettings) SMTPServer() string {
	return e.smtpServer
}

// SMTPUsername returns the applications SMTP username
func (e *emailSettings) SMTPUsername() string {
	return e.smtpUsername
}

// SMTPPassword returns the applications SMTP password
func (e *emailSettings) SMTPPassword() string {
	return e.smtpPassword
}

// SMTPPort returns the applications SMTP port
func (e *emailSettings) SMTPPort() int {
	return e.smtpPort
}

// ContactEmail returns the applications contact email
func (e *emailSettings) ContactEmail() string {
	return e.contactEmail
}
