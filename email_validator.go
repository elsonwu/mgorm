package mgorm

import (
	"errors"
	"reflect"
	"regexp"
)

const EXP_EMAIL = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

func EmailValidator(fieldValue reflect.Value, fieldType reflect.StructField) error {
	ok, err := regexp.MatchString(EXP_EMAIL, fieldValue.String())
	if nil != err {
		return err
	}

	if !ok {
		return errors.New(fieldType.Name + " is not email")
	}

	return nil
}
