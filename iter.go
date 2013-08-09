package mgorm

import (
	"labix.org/v2/mgo"
)

type Iter struct {
	iter  *mgo.Iter
	query *Query
}

func (self *Iter) Next(model IModel) bool {
	b := self.iter.Next(model)
	model.Init()
	model.AfterFind()

	if !b {
		self.iter.Close()
	}

	return b
}

func (self *Iter) All(models interface{}) {
	self.iter.All(models)
}

func (self *Iter) Close() error {
	return self.iter.Close()
}
