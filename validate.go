package mgorm

import (
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
		if reflect.Struct != field.Type.Kind() || "Model" == field.Name {
			continue
		}

		reFiled := reflect.New(refValue.Field(i).Type())
		if !reFiled.MethodByName("Validate").IsValid() {
			continue
		}

		res := reFiled.MethodByName("Validate").Call([]reflect.Value{})
		if res[0].Bool() {
			return false
		}
	}

	return true
}
