package model

import (
	"github.com/elsonwu/restapi/model/attr"
	"labix.org/v2/mgo/bson"
)

type Criteria struct {
	limit      attr.Int
	offset     attr.Int
	conditions attr.Map
}

func (self *Criteria) GetLimit() int {
	return int(self.limit)
}

func (self *Criteria) SetLimit(limit int) {
	self.limit = attr.Int(limit)
}

func (self *Criteria) GetOffset() int {
	return int(self.offset)
}

func (self *Criteria) SetOffset(offset int) {
	self.offset = attr.Int(offset)
}

func (self *Criteria) GetConditions() bson.M {
	return bson.M(self.conditions)
}

func (self *Criteria) SetConditions(conditions attr.Map) {
	self.conditions = conditions
}

func (self *Criteria) AddCondition(condition attr.Map) {
	for k, v := range condition {
		if _, ok := self.conditions[k]; ok {
			//@todo merge if some case
			self.conditions[k] = v
		} else {
			self.conditions[k] = v
		}
	}
}
