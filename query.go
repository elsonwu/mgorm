package model

import (
	"labix.org/v2/mgo"
)

type Query struct {
	query *mgo.Query
}

func (self *Query) GetQuery() *mgo.Query {
	return self.query
}

func (self *Query) SetQuery(query *mgo.Query) {
	self.query = query
}

func (self *Query) One(model IModel) {
	self.query.One(model)
}

func (self *Query) Iter() *Iter {
	iter := new(Iter)
	iter.SetIter(self.query.Iter())
	return iter
}
