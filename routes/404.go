package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-playground/log"
)

// get404Handler function renders a 404 page when a page
// can not be found on the server
func (a *app) get404Handler(w http.ResponseWriter, r *http.Request) {

	t := a.Translator(r)

	trans := struct {
		Title   string
		Message template.HTML
	}{
		Title:   "404 " + t.T("404-not-found"),
		Message: template.HTML(t.T("404-not-found-msg", fmt.Sprintf("<a href=\"/\">%s</a>", t.T("404-home")))),
	}

	err := a.ExecuteTemplate(w, "404", trans, nil)
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing 404 Template")
	}
}
