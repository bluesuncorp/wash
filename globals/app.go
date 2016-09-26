package globals

import (
	"html/template"
	"io"
	"net/http"

	"github.com/bluesuncorp/wash/env"
	"github.com/go-playground/pure"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

// App contains the application level context
type App struct {
	buffer    ByteBuffer
	ut        *ut.UniversalTranslator
	email     EmailSettings
	validate  *validator.Validate
	templates Templates
}

// NewApp creates a new App object instance
func NewApp(cfg *env.Config, ut *ut.UniversalTranslator, validate *validator.Validate, templates *template.Template) *App {

	return &App{
		buffer:    newByteBuffer(),
		email:     newEmail(cfg.SMTPServer, cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPPort, cfg.SupportEmail),
		ut:        ut,
		validate:  validate,
		templates: newTemplates(templates),
	}
}

// ByteBuffer returns the global buffer instance
func (a *App) ByteBuffer() ByteBuffer {
	return a.buffer
}

//Email returns the global email settings information
func (a *App) Email() EmailSettings {
	return a.email
}

// Templates returns the html template object
func (a *App) Templates() Templates {
	return a.templates
}

// Validator returns the applications validator
func (a *App) Validator() *validator.Validate {
	return a.validate
}

// ExecuteTemplate calls the regular Tempaltes ExecuteTemplate and automatically uses the c.Response() object
// as the io.Writer, if you wish to write to somethinf else call c.Templates().ExecuteTemplate()
func (a *App) ExecuteTemplate(w io.Writer, name string, translations interface{}, data interface{}) error {
	return a.templates.(*templates).executeTemplate(w, name, &templateData{App: a, Trans: translations, Data: data})
}

// Translator returns the applications Translator object
func (a *App) Translator(r *http.Request) Translator {

	// try really hard to find locale in this order ( feel free to comment out any or all that do not apply to your app):
	// - query param
	// - cookie
	// - http 'Accept-Header'

	var trans ut.Translator
	var found bool

	locale := r.URL.Query().Get("locale")
	if len(locale) > 0 {

		if trans, found = a.ut.GetTranslator(locale); found {
			goto END
		}

		// can't redirect, not in middleware or handlers...
		// perhaps add a translator middleware check and if not set
		// and redirect?

		// as of now this app just uses the fallback.
	}

	// Try and get from cookie value, should only be set on users that we could not determine the language for
	if cookie, err := r.Cookie("locale"); err == nil {

		if trans, found = a.ut.GetTranslator(cookie.Value); found {
			goto END
		}

		// can't redirect, not in middleware or handlers...
		// perhaps add a translator middleware check and if not set
		// and redirect?

		// as of now this app just uses the fallback.
	}

	// if not found, that's ok will use fallback so ignoring return value
	trans, _ = a.ut.FindTranslator(pure.AcceptedLanguages(r)...)

END:
	return newTranslator(trans)
}
