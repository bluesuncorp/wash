package globals

import (
	"bytes"
	"net/http"

	"github.com/go-playground/lars"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

// Context is the http level context
type Context struct {
	*lars.Ctx
	app       *App
	templates *Templates
	ut        *ut.UniversalTranslator
	buff      *ByteBuffer
}

var _ lars.Context = new(Context)

// NewContext returns a new http level Context instance
func NewContext(l *lars.LARS, templates *Templates, buff *ByteBuffer, ut *ut.UniversalTranslator, e *EmailSettings, validator *validator.Validate) lars.ContextFunc {

	return func(l *lars.LARS) lars.Context {
		return &Context{
			Ctx:       lars.NewContext(l),
			app:       NewApp(buff, ut, e, validator),
			templates: templates,
			ut:        ut,
			buff:      buff,
		}
	}
}

// App returns the application level context registered with Context
func (c *Context) App() *App {
	return c.app
}

// Templates returns the html template object
func (c *Context) Templates() *Templates {
	return c.templates
}

// ExecuteTemplate calls the regular Tempaltes ExecuteTemplate and automatically uses the c.Response() object
// as the io.Writer, if you wish to write to somethinf else call c.Templates().ExecuteTemplate()
func (c *Context) ExecuteTemplate(name string, translations interface{}, data interface{}) error {
	return c.templates.executeTemplate(c.Response(), name, &templateData{Ctx: c, Trans: translations, Data: data})
}

// ExecuteAndReturnTemplate calls the regular Tempaltes ExecuteTemplate and automatically uses the c.Response() object
// as the io.Writer, if you wish to write to somethinf else call c.Templates().ExecuteTemplate()
func (c *Context) ExecuteAndReturnTemplate(name string, translations interface{}, data interface{}) ([]byte, error) {
	b := bytes.NewBuffer(c.buff.Get())
	err := c.templates.executeTemplate(b, name, &templateData{Ctx: c, Trans: translations, Data: data})
	return b.Bytes(), err
}

// RequestStart is called just as the request is about to start, for setting init variables and alike
func (c *Context) RequestStart(w http.ResponseWriter, r *http.Request) {

	c.Ctx.RequestStart(w, r)

	// try really hard to find locale in this order ( feel free to comment out any or all that do not apply to your app):
	// - query param
	// - cookie
	// - http 'Accept-Header'

	var trans ut.Translator
	var found bool

	locale := c.Ctx.QueryParams()["locale"]
	if len(locale) > 0 {

		if trans, found = c.ut.GetTranslator(locale[0]); found {
			goto END
		}

		// can't redirect, not in middleware or handlers...
		// perhaps add a translator middleware check and if not set
		// and redirect?

		// as of now this app just uses the fallback.
	}

	// Try and get from cookie value, should only be set on users that we could not determine the language for
	if cookie, err := r.Cookie("locale"); err == nil {

		if trans, found = c.ut.GetTranslator(cookie.Value); found {
			goto END
		}

		// can't redirect, not in middleware or handlers...
		// perhaps add a translator middleware check and if not set
		// and redirect?

		// as of now this app just uses the fallback.
	}

	// if not found, that's ok will use fallback so ignoring return value
	trans, _ = c.ut.FindTranslator(c.Ctx.AcceptedLanguages(false)...)

END:
	c.app.trans = newTranslator(trans)
}

// CastContext function casts lars.Context to our Context prior to calling handlers
func CastContext(c lars.Context, handler lars.Handler) {
	handler.(func(*Context))(c.(*Context))
}
