package param_verify

import (
	"gopkg.in/go-playground/validator.v9"
)

func NameValid(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "admin"  {
		return false
	}
	return true
}
