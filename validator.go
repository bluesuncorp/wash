package main

import "gopkg.in/go-playground/validator.v9"

func initValidator() *validator.Validate {
	return validator.New()
}
