package globals

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/log"
)

// Translator wraps ut.Translator in order to wrap and handle errors transparently
type Translator interface {
	T(key interface{}, params ...string) string
	C(key interface{}, num float64, digits uint64, param string) string
	O(key interface{}, num float64, digits uint64, param string) string
	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string
	Base() ut.Translator
}

type translator struct {
	trans ut.Translator
}

var _ Translator = new(translator)

func newTranslator(trans ut.Translator) Translator {
	return &translator{trans: trans}
}

// Base returns the underlying Translator that is wrapped by the globals Translator.
// this is needed for validator translations
func (t *translator) Base() ut.Translator {
	return t.trans
}

// creates the translation for the locale given the 'key' and params passed in
func (t *translator) T(key interface{}, params ...string) string {

	s, err := t.trans.T(key, params...)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("error", err)).Warn("issue translating")
	}

	return s
}

// creates the cardinal translation for the locale given the 'key', 'num' and 'digit' arguments
//  and param passed in
func (t *translator) C(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("error", err)).Warn("issue translating cardinal")
	}

	return s
}

// creates the ordinal translation for the locale given the 'key', 'num' and 'digit' arguments
// and param passed in
func (t *translator) O(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.O(key, num, digits, param)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("error", err)).Warn("issue translating ordinal")
	}

	return s
}

//  creates the range translation for the locale given the 'key', 'num1', 'digit1', 'num2' and
//  'digit2' arguments and 'param1' and 'param2' passed in
func (t *translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string {

	s, err := t.trans.R(key, num1, digits1, num2, digits2, param1, param2)
	if err != nil {
		log.StackTrace().WithFields(
			log.F("key", key),
			log.F("error", err)).Warn("issue translating range")
	}

	return s
}
