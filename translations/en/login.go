package english

import (
	"github.com/bluesuncorp/wash/helpers/trans"
	"github.com/go-playground/universal-translator"
)

func login(en ut.Translator) {

	trans.Add(en, "login-login", "Login", false)
	trans.Add(en, "login-welcome", "Welcome {0}", false)
}
