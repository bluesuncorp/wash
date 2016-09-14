package english

import (
	"github.com/bluesuncorp/wash/helpers/trans"
	"github.com/go-playground/universal-translator"
)

func jsRequired(en ut.Translator) {

	trans.Add(en, "js-req-required", "Javascript Required", false)
	trans.Add(en, "js-req-required-msg", "This site requires Javascript to function correctly", false)
}
