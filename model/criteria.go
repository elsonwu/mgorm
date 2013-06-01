package model

import (
	"github.com/elsonwu/restapi/model/attr"
)

type Criteria struct {
	limit      attr.Int
	offset     attr.Int
	conditions attr.Map
}

func (self *Criteria) GetLimit() int {
	return self.limit.Get()
}

func (self *Criteria) SetLimit(limit int) {
	self.limit = attr.Int(limit)
}

func (self *Criteria) GetOffset() int {
	return self.offset.Get()
}

func (self *Criteria) SetOffset(offset int) {
	self.offset = attr.Int(offset)
}

func (self *Criteria) GetConditions() attr.Map {
	return self.conditions
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
