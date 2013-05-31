package attr

import (
	"regexp"
)

type Email string

func (self *Email) Validate() (boolean bool) {
	boolean, _ = regexp.MatchString(".*@.*", string(*self))
	return
}

func (self *Email) Set(v string) {
	*self = Email(v)
}
