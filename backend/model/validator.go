package model

import "gopkg.in/go-playground/validator.v9"

var Validator = &customValidator{validator: validator.New()}

type customValidator struct {
	validator *validator.Validate
}

func (v *customValidator) Validate(model interface{}) error {
	return v.validator.Struct(model)
}
