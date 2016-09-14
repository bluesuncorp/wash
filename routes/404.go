package routes

import (
	"fmt"
	"html/template"

	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

// Get404Handler function renders a 404 page when a page
// can not be found on the server
func Get404Handler(c *globals.Context) {

	t := c.App().Translator()

	trans := struct {
		Title   string
		Message template.HTML
	}{
		Title:   "404 " + t.T("404-not-found"),
		Message: template.HTML(t.T("404-not-found-msg", fmt.Sprintf("<a href=\"/\">%s</a>", t.T("404-home")))),
	}

	err := c.ExecuteTemplate("404", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing 404 Template")
	}
}
