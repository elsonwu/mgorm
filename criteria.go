package mgorm

import ()

func NewCriteria() ICriteria {
	criteria := new(Criteria)
	criteria.conditions = Map{}
	return criteria
}

const CriteriaSortDesc int = -1
const CriteriaSortAsc int = 1

type ICriteria interface {
	AddSort(field string, sort int) ICriteria
	SetSort(sort map[string]int) ICriteria
	GetSort() map[string]int
	SetSelect(selects []string) ICriteria
	GetSelect() []string
	GetLimit() int
	SetLimit(limit int) ICriteria
	GetOffset() int
	SetOffset(offset int) ICriteria
	GetConditions() Map
	SetConditions(conditions Map) ICriteria
	AddCond(field, opt string, value interface{}) ICriteria
}

type Criteria struct {
	selects    []string
	limit      int
	offset     int
	conditions Map
	sort       map[string]int
}

func (self *Criteria) AddSort(field string, sort int) ICriteria {
	if nil == self.sort {
		self.sort = map[string]int{}
	}

	self.sort[field] = sort
	return self
}

func (self *Criteria) SetSort(sort map[string]int) ICriteria {
	self.sort = sort
	return self
}

func (self *Criteria) GetSort() map[string]int {
	return self.sort
}

func (self *Criteria) SetSelect(selects []string) ICriteria {
	self.selects = selects
	return self
}

func (self *Criteria) GetSelect() []string {
	return self.selects
}

func (self *Criteria) GetLimit() int {
	return self.limit
}

func (self *Criteria) SetLimit(limit int) ICriteria {
	self.limit = limit
	return self
}

func (self *Criteria) GetOffset() int {
	return self.offset
}

func (self *Criteria) SetOffset(offset int) ICriteria {
	self.offset = offset
	return self
}

func (self *Criteria) GetConditions() Map {
	return self.conditions
}

func (self *Criteria) SetConditions(conditions Map) ICriteria {
	self.conditions = conditions
	return self
}

//add more opt later when it does need.
func (self *Criteria) AddCond(field, opt string, value interface{}) ICriteria {
	if opt == "==" {
		self.conditions[field] = value
	} else if opt == "!=" {
		self.conditions[field] = Map{"$ne": value}
	} else if opt == "<" {
		self.conditions[field] = Map{"$lt": value}
	} else if opt == "<=" {
		self.conditions[field] = Map{"$lte": value}
	} else if opt == ">" {
		self.conditions[field] = Map{"$gt": value}
	} else if opt == ">=" {
		self.conditions[field] = Map{"$gte": value}
	} else if opt == "in" {
		self.conditions[field] = Map{"$in": value}
	} else if opt == "nin" {
		self.conditions[field] = Map{"$nin": value}
	} else if opt == "size" {
		self.conditions[field] = Map{"$size": value}
	} else if opt == "all" {
		self.conditions[field] = Map{"$all": value}
	} else if opt == "where" {
		self.conditions[field] = Map{"$where": value}
	} else if opt == "type" {
		self.conditions[field] = Map{"$type": value}
	} else if opt == "exists" {
		self.conditions[field] = Map{"$exists": value}
	} else if opt == "or" {
		if v, ok := self.conditions["$or"]; ok && v != nil {
			if or, ok := v.([]Map); ok {
				or = append(or, Map{field: value})
				self.conditions["$or"] = or
			}
		} else {
			or := make([]Map, 1)
			or[0] = Map{field: value}
			self.conditions["$or"] = or
		}
	}

	return self
}
