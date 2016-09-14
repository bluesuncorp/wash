package routes

import (
	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

func getLogin(c *globals.Context) {

	t := c.App().Translator()

	trans := struct {
		Title   string
		Message string
	}{
		Title:   t.T("login-login"),
		Message: t.T("login-welcome", "Joeybloggs"),
	}

	err := c.ExecuteTemplate("login", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing login Template")
	}
}
