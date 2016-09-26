package routes

import (
	"net/http"

	"github.com/go-playground/log"
)

func (a *app) getRoot(w http.ResponseWriter, r *http.Request) {

	// redirect to login or specific homepage....
	_, err := w.Write([]byte("Redirect ME or make me your HOMEPAGE!"))
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing root Template")
	}
}
