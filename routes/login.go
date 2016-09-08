package routes

import (
	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

func getLogin(c *globals.Context) {

	t := c.App().Translator()

	visitor := t.O("numsuffix", 3, 0, "3")
	days := t.C("days", 2, 0, "2")
	final := t.T("testonly", "Joeybloggs", visitor, days)

	s := struct {
		Title   string
		Message string
	}{
		Title:   "Login",
		Message: final,
	}

	err := c.ExecuteTemplate("login", s)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing login Template")
	}
}
