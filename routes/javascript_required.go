package routes

import (
	"net/http"

	"github.com/go-playground/log"
)

// getJavascriptRequiredHandler will process the JavaScript required page explaining that this
// site requires JavaScript
func (a *app) getJavascriptRequiredHandler(w http.ResponseWriter, r *http.Request) {

	t := a.Translator(r)

	trans := struct {
		Title   string
		Message string
	}{
		Title:   t.T("js-req-required"),
		Message: t.T("js-req-required-msg"),
	}

	err := a.ExecuteTemplate(w, "javascript-required", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing Javascript Required Template")
	}
}
