package model

import (
	// "fmt"
	"github.com/elsonwu/restapi/model/attr"
)

func NewCriteria() Criteria {
	criteria := Criteria{}
	criteria.conditions = attr.Map{}
	return criteria
}

type Criteria struct {
	limit      attr.Int
	offset     attr.Int
	sort       attr.Map
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

//add more opt later when it does need.
func (self *Criteria) AddCondition(field, opt string, value interface{}) {
	if opt == "==" {
		self.conditions[field] = value
	} else if opt == "!=" {
		self.conditions[field] = attr.Map{"$ne": value}
	} else if opt == "<" {
		self.conditions[field] = attr.Map{"$lt": value}
	} else if opt == "<=" {
		self.conditions[field] = attr.Map{"$lte": value}
	} else if opt == ">" {
		self.conditions[field] = attr.Map{"$gt": value}
	} else if opt == ">=" {
		self.conditions[field] = attr.Map{"$gte": value}
	} else if opt == "in" {
		self.conditions[field] = attr.Map{"$in": value}
	} else if opt == "nin" {
		self.conditions[field] = attr.Map{"$nin": value}
	} else if opt == "size" {
		self.conditions[field] = attr.Map{"$size": value}
	} else if opt == "all" {
		self.conditions[field] = attr.Map{"$all": value}
	} else if opt == "where" {
		self.conditions[field] = attr.Map{"$where": value}
	} else if opt == "type" {
		self.conditions[field] = attr.Map{"$type": value}
	} else if opt == "exists" {
		self.conditions[field] = attr.Map{"$exists": value}
	} else if opt == "or" {
		if v, ok := self.conditions["$or"]; ok && v != nil {
			if or, ok := v.([]attr.Map); ok {
				or = append(or, attr.Map{field: value})
				self.conditions["$or"] = or
			}
		} else {
			or := make([]attr.Map, 1)
			or[0] = attr.Map{field: value}
			self.conditions["$or"] = or
		}
	}
}
