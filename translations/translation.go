package translations

import (
	"github.com/bluesuncorp/wash/translations/en"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

// Initialize initializes and returns the UniversalTranslator instance for the application
func Initialize(validate *validator.Validate) *ut.UniversalTranslator {

	// initialize translator
	en := en.New()
	uni := ut.New(en, en)

	// initialize translations
	english.Init(uni, validate)

	return uni
}
