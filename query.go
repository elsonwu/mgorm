package mgorm

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Query struct {
	query *mgo.Query
	iter  *mgo.Iter
}

func (self *Query) Iter() *mgo.Iter {
	if nil == self.iter {
		self.iter = self.query.Iter()
	}

	return self.iter
}

func (self *Query) Query() *mgo.Query {
	return self.query
}

func (self *Query) One(model IModel) {
	err := self.query.One(model)
	if nil == err {
		model.AfterFind()
	}
}

func (self *Query) Next(model IModel) bool {
	b := self.Iter().Next(model)
	model.AfterFind()

	if !b {
		self.Iter().Close()
	}

	return b
}

func (self *Query) All(models interface{}) {
	self.Iter().All(models)
}

func (self *Query) Close() error {
	return self.Iter().Close()
}

func (self *Query) Count() int {
	count, err := self.query.Count()
	if nil != err {
		fmt.Println("count err:", err)
	}

	return count
}
