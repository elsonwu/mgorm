package mgorm

import (
	// "fmt"
	"reflect"
)

type IValidator interface {
	Validate() bool
}

func NewValidator(embeddedModel IErrorHandler) IValidator {
	validater := new(Validator)
	validater.embeddedModel = embeddedModel
	return validater
}

type ValidateFn func(fieldValue reflect.Value, fieldType reflect.StructField) error

type Validator struct {
	embeddedModel IErrorHandler
}

func (self *Validator) Validate() bool {
	refType := reflect.TypeOf(self.embeddedModel)
	refValue := reflect.ValueOf(self.embeddedModel)

	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	numField := refType.NumField()

	var hasError bool
	for i := 0; i < numField; i++ {
		fieldType := refType.Field(i)
		fieldValue := refValue.Field(i)

		var err error
		if reflect.String == fieldType.Type.Kind() {
			tag := fieldType.Tag.Get("rules")
			switch tag {
			case "email":
				err = Validate(fieldValue, fieldType, EmailValidator)
			case "url":
				err = Validate(fieldValue, fieldType, UrlValidator)
			default:
				//do nothing
			}

			if nil != err {
				self.embeddedModel.AddError(err.Error())
				hasError = true
			}
		}

		if reflect.Ptr != fieldType.Type.Kind() {
			continue
		}

		if v, ok := fieldValue.Interface().(IValidator); ok {
			if !fieldValue.IsNil() && !v.Validate() {
				hasError = true
			}
		}

		if v, ok := fieldValue.Interface().(IErrorHandler); ok {
			if !fieldValue.IsNil() && !NewValidator(v).Validate() {
				hasError = true
			}
		}
	}

	return hasError
}

func Validate(fieldValue reflect.Value, fieldType reflect.StructField, fn ValidateFn) error {
	return fn(fieldValue, fieldType)
}
