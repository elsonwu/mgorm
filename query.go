package mgorm

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Query struct {
	query *mgo.Query
	iter  *Iter
}

func (self *Query) Query() *mgo.Query {
	return self.query
}

func (self *Query) One(model IModel) {
	self.query.One(model)
}

func (self *Query) Iter() *Iter {
	if nil == self.iter {
		self.iter = new(Iter)
		self.iter.query = self
		self.iter.iter = self.query.Iter()
	}

	return self.iter
}

func (self *Query) Count() int {
	count, err := self.query.Count()
	if nil != err {
		fmt.Println("count err:", err)
	}

	return count
}
