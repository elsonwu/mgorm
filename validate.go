package mgorm

import (
	// "fmt"
	"reflect"
)

type IValidater interface {
	Validate() bool
}

func NewValidater(errorHandler IErrorHandler) *Validater {
	validater := new(Validater)
	validater.errorHandler = errorHandler
	return validater
}

type Validater struct {
	errorHandler IErrorHandler
}

func (self *Validater) Validate() bool {
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

		if v, ok := refValue.Field(i).Interface().(IValidater); ok {
			if !v.Validate() {
				return false
			}
		}
	}

	return true
}
