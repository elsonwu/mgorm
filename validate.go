package mgorm

import (
	// "fmt"
	"reflect"
)

type IValidator interface {
	Validate() bool
}

func NewValidator(errorHandler IErrorHandler) IValidator {
	validater := new(Validator)
	validater.errorHandler = errorHandler
	return validater
}

type Validator struct {
	errorHandler IErrorHandler
}

func (self *Validator) Validate() bool {
	refType := reflect.TypeOf(self.errorHandler)
	refValue := reflect.ValueOf(self.errorHandler)

	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	numField := refType.NumField()
	for i := 0; i < numField; i++ {
		field := refType.Field(i)

		if reflect.Ptr != field.Type.Kind() {
			continue
		}

		if v, ok := refValue.Field(i).Interface().(IValidator); ok {
			if !v.Validate() {
				return false
			}
		}
	}

	return true
}
