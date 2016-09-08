package routes

import (
	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

// Get404Handler function renders a 404 page when a page
// can not be found on the server
func Get404Handler(c *globals.Context) {

	s := struct {
		Title string
	}{
		Title: "404 Not Found",
	}

	err := c.ExecuteTemplate("404", s)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing 404 Template")
	}
}
