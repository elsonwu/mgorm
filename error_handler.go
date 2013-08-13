package mgorm

import (
	"errors"
)

type IErrorHandler interface {
	HasErrors() bool
	AddError(err string)
	GetErrors() []error
	ClearErrors()
}

type ErrorHandler struct {
	errors []error `bson:",inline" json:"-"`
}

func (self *ErrorHandler) AddError(err string) {
	self.errors = append(self.errors, errors.New(err))
}

func (self *ErrorHandler) GetErrors() []error {
	return self.errors
}

func (self *ErrorHandler) ClearErrors() {
	self.errors = []error{}
}

func (self *ErrorHandler) HasErrors() bool {
	return 0 < len(self.errors)
}
