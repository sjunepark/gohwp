package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func initValidator() {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
	}
}

func ValidateStruct(i interface{}) error {
	initValidator()
	return validate.Struct(i)
}
