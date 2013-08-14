package mgorm

import (
	// "fmt"
	"reflect"
)

type IValidator interface {
	Validate() bool
}

type ValidateFn func(fieldValue reflect.Value, fieldType reflect.StructField) error

func Validate(model IEmbeddedModel) bool {
	refType := reflect.TypeOf(model)
	refValue := reflect.ValueOf(model)

	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	numField := refType.NumField()

	isOK := true
	for i := 0; i < numField; i++ {
		fieldType := refType.Field(i)
		fieldValue := refValue.Field(i)

		var err error
		if reflect.String == fieldType.Type.Kind() {
			tag := fieldType.Tag.Get("rules")
			switch tag {
			case "email":
				err = fieldValidate(fieldValue, fieldType, EmailValidator)
			case "url":
				err = fieldValidate(fieldValue, fieldType, UrlValidator)
			default:
				//do nothing
			}

			if nil != err {
				model.AddError(err.Error())
				isOK = false
			}
		}

		if reflect.Ptr != fieldType.Type.Kind() && reflect.Struct != fieldType.Type.Kind() {
			continue
		}

		if reflect.Struct == fieldType.Type.Kind() {
			fieldValue = fieldValue.Addr()
		}

		if v, ok := fieldValue.Interface().(IValidator); ok {
			if !v.Validate() {
				isOK = false
			}
		}

		if v, ok := fieldValue.Interface().(IEmbeddedModel); ok {
			if !Validate(v) {
				isOK = false
			}
		}
	}

	return isOK
}

func fieldValidate(fieldValue reflect.Value, fieldType reflect.StructField, fn ValidateFn) error {
	return fn(fieldValue, fieldType)
}
