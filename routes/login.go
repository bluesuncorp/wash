package routes

import (
	"net/http"

	"github.com/go-playground/log"
)

func (a *app) getLogin(w http.ResponseWriter, r *http.Request) {

	t := a.Translator(r)

	trans := struct {
		Title   string
		Message string
	}{
		Title:   t.T("login-login"),
		Message: t.T("login-welcome", "Joeybloggs"),
	}

	err := a.ExecuteTemplate(w, "login", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing login Template")
	}
}
