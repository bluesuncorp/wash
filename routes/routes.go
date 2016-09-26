package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bluesuncorp/wash/env"
	"github.com/bluesuncorp/wash/globals"
	"github.com/bluesuncorp/wash/middleware"
	"github.com/go-playground/log"
	"github.com/go-playground/pure"
	mw "github.com/go-playground/pure/middleware"
)

const (
	appPath = "$GOPATH/src/github.com/bluesuncorp/wash"
)

type app struct {
	*globals.App
}

var a *app

// Initialize initializes and return the http routes
func Initialize(p *pure.Mux, globs *globals.App, cfg *env.Config) (redirect *pure.Mux) {

	a = &app{App: globs}

	log.Info("Initializing Routes ...")

	p.Use(middleware.LoggingAndRecovery, mw.Gzip, middleware.Security)

	// different middlewae for 404 as once we start adding CSRF and auth middlewae, thye don;t really
	// apply to the 404 page
	p.Register404(a.get404Handler, middleware.LoggingAndRecovery, mw.Gzip, middleware.Security)

	p.Get("/javascript-required", a.getJavascriptRequiredHandler)
	p.Get("/", a.getRoot)
	p.Get("/login", a.getLogin)

	initAssets(p)

	if cfg.IsProduction {
		redirect = setupRedirect(cfg)
	}

	return
}

func initAssets(r pure.IRouteGroup) {

	fs := http.FileServer(http.Dir("assets"))

	a := r.Group("/assets/", nil)
	a.Use(mw.Gzip, middleware.Security)
	a.Use(func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			if strings.LastIndex(r.URL.Path, ".") == -1 {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next(w, r)
		}
	})

	// also add authentication?
	a.Get("*", http.StripPrefix("/assets", fs).ServeHTTP)

	i := r.Group("", nil)
	i.Use(middleware.LoggingAndRecovery, mw.Gzip, middleware.Security)

	i.Get("/favicon.ico", fs.ServeHTTP)
	i.Get("/robots.txt", fs.ServeHTTP)
}

func setupRedirect(cfg *env.Config) (redirect *pure.Mux) {

	redirect = pure.New()
	redirect.Use(middleware.LoggingAndRecovery)
	redirect.Get("*", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+":"+strconv.Itoa(cfg.RedirectPort)+r.URL.String(), http.StatusMovedPermanently)
	})
	return
}
