package trans

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/log"
	"github.com/go-playground/universal-translator"
)

// Add add normal translation and wraps error
func Add(trans ut.Translator, key interface{}, text string, override bool) {

	err := trans.Add(key, text, override)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("text", text),
			log.F("error", err),
		).Fatal("issue adding translation")
	}
}

// AddRange adds a range translation and wraps error
func AddRange(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) {

	err := trans.AddRange(key, text, rule, override)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("text", text),
			log.F("rule", rule),
			log.F("error", err),
		).Fatal("issue adding range translation")
	}
}

// AddCardinal adds a cardinal translation and wraps error
func AddCardinal(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) {

	err := trans.AddCardinal(key, text, rule, override)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("text", text),
			log.F("rule", rule),
			log.F("error", err),
		).Fatal("issue adding cardinal translation")
	}
}

// AddOrdinal adds an ordinal translation and wraps error
func AddOrdinal(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) {

	err := trans.AddOrdinal(key, text, rule, override)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("text", text),
			log.F("rule", rule),
			log.F("error", err),
		).Fatal("issue adding ordinal translation")
	}
}
