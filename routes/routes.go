package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bluesuncorp/wash/env"
	"github.com/bluesuncorp/wash/middleware"
	"github.com/go-playground/lars"
	mw "github.com/go-playground/lars/middleware"
	"github.com/go-playground/log"
)

const (
	favicon     = "/favicon.ico"
	realFavicon = "/assets/images/favicon.ico"
	robots      = "/robots.txt"
	realRobots  = "/assets/robots.txt"
)

// Initialize initializes and return the http routes
func Initialize(l *lars.LARS, cfg *env.Config) (redirect *lars.LARS) {

	log.Info("Initializing Routes ...")

	l.Use(middleware.LoggingAndRecovery, mw.Gzip)

	l.Register404(Get404Handler)
	l.Get("/javascript-required", GetJavascriptRequiredHandler)

	l.Get("/", getRoot)
	l.Get("/login", getLogin)

	initAssets(l)

	if cfg.IsProduction {
		redirect = setupRedirect(cfg)
	}

	return
}

func initAssets(r lars.IRouteGroup) {

	fs := http.FileServer(http.Dir("assets"))

	a := r.Group("/assets/", nil)
	a.Use(middleware.Security, mw.Gzip)
	a.Use(func(c lars.Context) {

		if strings.LastIndex(c.Request().URL.Path, ".") == -1 {
			http.Error(c.Response(), http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		c.Next()
	})

	// also add authentication?
	a.Get("*", http.StripPrefix("/assets", fs))

	i := r.Group("", nil)
	i.Use(middleware.LoggingAndRecovery, middleware.Security, mw.Gzip)

	fav := func(c lars.Context) {
		c.Request().URL.Path = realFavicon
		c.Next()
		c.Request().URL.Path = favicon
	}

	robots := func(c lars.Context) {
		c.Request().URL.Path = realRobots
		c.Next()
		c.Request().URL.Path = robots
	}

	i.Get("/favicon.ico", fav, fs)
	i.Get("/robots.txt", robots, fs)
}

func setupRedirect(cfg *env.Config) (redirect *lars.LARS) {

	redirect = lars.New()
	redirect.Use(middleware.LoggingAndRecovery)
	redirect.Get("*", func(c lars.Context) {
		req := c.Request()
		http.Redirect(c.Response(), req, "https://"+req.Host+":"+strconv.Itoa(cfg.RedirectPort)+req.URL.String(), http.StatusMovedPermanently)
	})
	return
}
