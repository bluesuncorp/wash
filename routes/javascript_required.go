package routes

import (
	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

// GetJavascriptRequiredHandler will process the JavaScript required page explaining that this
// site requires JavaScript
func GetJavascriptRequiredHandler(c *globals.Context) {

	t := c.App().Translator()

	trans := struct {
		Title   string
		Message string
	}{
		Title:   t.T("js-req-required"),
		Message: t.T("js-req-required-msg"),
	}

	err := c.ExecuteTemplate("javascript-required", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing Javascript Required Template")
	}
}
