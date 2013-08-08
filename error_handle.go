package model

type IErrorHandle interface {
	HasError() bool
	AddError(err string)
	GetErrors() []string
	ClearErrors()
}

type ErrorHandle struct {
	errors []string `bson:",omitempty" json:"-"`
}

func (self *ErrorHandle) AddError(err string) {
	self.errors = append(self.errors, err)
}

func (self *ErrorHandle) GetErrors() []string {
	return self.errors
}

func (self *ErrorHandle) ClearErrors() {
	self.errors = make([]string, 1)
}

func (self *ErrorHandle) HasError() bool {
	return !(self.errors == nil || 0 == len(self.errors))
}
