package routes

import (
	"net/http"

	"github.com/bluesuncorp/wash/globals"
	"github.com/go-playground/log"
)

func getRoot(c *globals.Context) {

	// redirect to login or specific homepage....
	err := c.Text(http.StatusOK, "Redirect ME or make me your HOMEPAGE!")
	if err != nil {
		log.WithFields(log.F("error", err)).Error("Issue Executing root Template")
	}
}
