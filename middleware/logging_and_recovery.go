package middleware

import (
	"log"
	"time"

	"github.com/go-playground/lars"
)

// LoggingAndRecovery Function for http requests
func LoggingAndRecovery(c lars.Context) {

	start := time.Now()

	c.Next()

	stop := time.Now()
	path := c.Request().URL.Path

	if path == "" {
		path = "/"
	}

	log.Printf("%s %d %s %s", c.Request().Method, c.Response().Status(), path, stop.Sub(start))
}
