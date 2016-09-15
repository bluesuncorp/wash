package english

import (
	"github.com/bluesuncorp/wash/helpers/trans"
	"github.com/go-playground/locales"
	"github.com/go-playground/log"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	valtrans "gopkg.in/go-playground/validator.v9/translations/en"
)

// Init initializes the english locale translations
func Init(uni *ut.UniversalTranslator, validate *validator.Validate) {

	locale := "en"

	en, found := uni.GetTranslator(locale)
	if !found {
		log.WithFields(log.F("locale", locale)).Fatal("Translation not found")
	}

	// validator translations & Overrides
	err := valtrans.RegisterDefaultTranslations(validate, en)
	if err != nil {
		log.WithFields(log.F("error", err)).Alert("Error adding default translations!")
	}

	// cardinals
	trans.AddCardinal(en, "days", "{0} day", locales.PluralRuleOne, false)
	trans.AddCardinal(en, "days", "{0} days", locales.PluralRuleOther, false)

	// ordinals
	trans.AddOrdinal(en, "numsuffix", "{0}st", locales.PluralRuleOne, false)
	trans.AddOrdinal(en, "numsuffix", "{0}nd", locales.PluralRuleTwo, false)
	trans.AddOrdinal(en, "numsuffix", "{0}rd", locales.PluralRuleFew, false)
	trans.AddOrdinal(en, "numsuffix", "{0}th", locales.PluralRuleOther, false)

	// ranges
	trans.AddRange(en, "dayrange", "{0}-{1} days", locales.PluralRuleOther, false)

	// actual message translations
	login(en)
	notFound(en)
	jsRequired(en)

	if err := en.VerifyTranslations(); err != nil {
		log.WithFields(log.F("error", err)).Fatal("Missing Translations!")
	}
}
