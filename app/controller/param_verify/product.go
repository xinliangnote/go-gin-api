package param_verify

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func NameValid (
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		if s == "admin" {
			return false
		}
	}
	return true
}
