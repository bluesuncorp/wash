package globals

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

// NewApp creates a new App object instance
func NewApp(buff *ByteBuffer, ut *ut.UniversalTranslator, e *EmailSettings, validator *validator.Validate) *App {

	return &App{
		buffer:    buff,
		email:     e,
		ut:        ut,
		validator: validator,
	}
}

// App contains the application level context
type App struct {
	buffer    *ByteBuffer
	ut        *ut.UniversalTranslator
	trans     Translator
	email     *EmailSettings
	validator *validator.Validate
	// etc...
}

// ByteBuffer returns the global buffer instance
func (a *App) ByteBuffer() *ByteBuffer {
	return a.buffer
}

//Email returns the global email settings information
func (a *App) Email() *EmailSettings {
	return a.email
}

// Translator returns the applications Translator object
func (a *App) Translator() Translator {

	if a.trans == nil {
		a.trans = newTranslator(a.ut.GetFallback())
	}

	return a.trans
}
