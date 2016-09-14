package english

import (
	"github.com/bluesuncorp/wash/helpers/trans"
	"github.com/go-playground/universal-translator"
)

func notFound(en ut.Translator) {

	trans.Add(en, "404-home", "Home", false)
	trans.Add(en, "404-not-found", "Not Found", false)
	trans.Add(en, "404-not-found-msg", "The resource you are looking for has moved or no longer exists. {0}", false)
}
