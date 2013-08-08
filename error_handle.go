package model

import (
	"errors"
)

type IErrorHandle interface {
	HasError() bool
	AddError(err string)
	GetErrors() []error
	ClearErrors()
}

type ErrorHandle struct {
	errors []error `bson:",omitempty" json:"-"`
}

func (self *ErrorHandle) AddError(err string) {
	self.errors = append(self.errors, errors.New(err))
}

func (self *ErrorHandle) GetErrors() []error {
	return self.errors
}

func (self *ErrorHandle) ClearErrors() {
	self.errors = []error{}
}

func (self *ErrorHandle) HasError() bool {
	return !(self.errors == nil || 0 == len(self.errors))
}
