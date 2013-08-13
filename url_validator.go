package mgorm

import (
	"errors"
	"reflect"
	"regexp"
)

const EXP_URL = `^((http|https|ftp){1}\:\/\/)?([a-zA-Z0-9\-\.]+[a-zA-Z0-9]+\.(net|cn|co|hk|tw|com|edu|gov|us|int|mil|org|int|mil|vg|uk|idv|tk|se|nz|nu|nl|ms|jp|jobs|it|ind|gen|firm|in|gs|fr|fm|eu|es|de|bz|be|at|am|ag|mx|asia|ws|xxx|tv|cc|ca|mobi|me|biz|arpa|info|name|pro|aero|coop|museum|ly|eg|mk)(:[a-zA-Z0-9]*)?\/?([a-zA-Z0-9\-\._\?\'\/\\\+&amp;%\$#\=~])*)+$`

func UrlValidator(fieldValue reflect.Value, fieldType reflect.StructField) error {
	ok, err := regexp.MatchString(EXP_URL, fieldValue.String())
	if nil != err {
		return err
	}

	if !ok {
		return errors.New(fieldType.Name + " is not url")
	}

	return nil
}
