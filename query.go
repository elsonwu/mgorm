package mgorm

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Query struct {
	query *mgo.Query
	iter  *Iter
}

func (self *Query) GetQuery() *mgo.Query {
	return self.query
}

func (self *Query) SetQuery(query *mgo.Query) {
	self.query = query
	self.iter = new(Iter)
	self.iter.SetIter(self.query.Iter())
}

func (self *Query) One(model IModel) {
	self.query.One(model)
}

func (self *Query) Iter() *Iter {
	return self.iter
}

func (self *Query) Count() int {
	count, err := self.query.Count()
	if nil != err {
		fmt.Println("count err:", err)
	}

	return count
}
