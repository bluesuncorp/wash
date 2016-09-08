package routes

import (
	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

// GetJavascriptRequiredHandler will process the JavaScript required page explaining that this
// site requires JavaScript
func GetJavascriptRequiredHandler(c *globals.Context) {

	s := struct {
		Title string
	}{
		Title: "Javascript Required",
	}

	err := c.ExecuteTemplate("javascript-required", s)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing Javascript Required Template")
	}
}
