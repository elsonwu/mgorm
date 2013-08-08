package model

import (
	"github.com/elsonwu/mgorm/model/attr"
)

func NewCriteria() *Criteria {
	criteria := new(Criteria)
	criteria.conditions = attr.Map{}
	return criteria
}

const CriteriaSortDesc int = -1
const CriteriaSortAsc int = 1

type Criteria struct {
	selects    []string
	limit      int
	offset     int
	conditions attr.Map
	sort       map[string]int
}

func (self *Criteria) AddSort(field string, sort int) *Criteria {
	if nil == self.sort {
		self.sort = map[string]int{}
	}

	self.sort[field] = sort
	return self
}

func (self *Criteria) SetSort(sort map[string]int) *Criteria {
	self.sort = sort
	return self
}

func (self *Criteria) GetSort() map[string]int {
	return self.sort
}

func (self *Criteria) SetSelect(selects []string) *Criteria {
	self.selects = selects
	return self
}

func (self *Criteria) GetSelect() []string {
	return self.selects
}

func (self *Criteria) GetLimit() int {
	return self.limit
}

func (self *Criteria) SetLimit(limit int) *Criteria {
	self.limit = limit
	return self
}

func (self *Criteria) GetOffset() int {
	return self.offset
}

func (self *Criteria) SetOffset(offset int) *Criteria {
	self.offset = offset
	return self
}

func (self *Criteria) GetConditions() attr.Map {
	return self.conditions
}

func (self *Criteria) SetConditions(conditions attr.Map) *Criteria {
	self.conditions = conditions
	return self
}

//add more opt later when it does need.
func (self *Criteria) AddCond(field, opt string, value interface{}) *Criteria {
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

	return self
}
