package env

import (
	"github.com/caarlos0/env"
	"github.com/go-playground/log"
	"gopkg.in/go-playground/validator.v9"
)

// Config is the application configuration
// read from ENV vars
type Config struct {
	Domain       string `env:"APP_DOMAIN"         envDefault:"localhost"               validate:"required"`
	IsProduction bool   `env:"APP_IS_PRODUCTION"`
	AppPort      int    `env:"APP_APP_PORT"       envDefault:"3005"                    validate:"required"`
	RedirectPort int    `env:"APP_REDIRECT_PORT"  envDefault:"8080"                    validate:"required"`
	SMTPServer   string `env:"APP_SMTP_SERVER"    envDefault:"localhost"               validate:"required"`
	SMTPUsername string `env:"APP_SMTP_USERNAME"  envDefault:""`
	SMTPPassword string `env:"APP_SMTP_PASSWORD"`
	SMTPPort     int    `env:"APP_SMTP_PORT"      envDefault:"1025"                    validate:"required"`
	SupportEmail string `env:"APP_SUPPORT_EMAIL"  envDefault:""                        validate:"required"`
}

// Parse parses, validates and then returns the application
// configuration based on ENV variables
func Parse(val *validator.Validate) (cfg *Config, err error) {

	log.Info("Parsing ENV vars...")
	defer log.Info("Done Parsing ENV vars")

	cfg = &Config{}

	if err = env.Parse(cfg); err != nil {
		log.WithFields(log.F("error", err)).Warn("Errors Parsing Configuration")
	}

	err = val.Struct(cfg)

	return
}
