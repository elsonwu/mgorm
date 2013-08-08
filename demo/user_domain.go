package main

import (
	// "fmt"
	"github.com/elsonwu/mgorm/model"
	"strconv"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

type UserDomain struct {
	owner             *User
	model.ErrorHandle `bson:",inline" json:"-"`
	Base              string `bson:"base" json:"base"`
	Extra             int    `bson:"extra" json:"extra"`
	Domain            string `bson:"domain" json:"domain"`
}

func (self *UserDomain) SetOwner(user *User) {
	self.owner = user
}

func (self *UserDomain) initDomain() bool {
	if self.owner.IsNew() {
		criteria := model.NewCriteria()
		criteria.AddCond("domain.base", "==", self.Base)
		criteria.AddSort("domain.domain", model.CriteriaSortDesc)
		user := new(User)
		query := model.FindAll(user, criteria)
		query.One(user)

		if user.Domain.Extra == 0 {
			self.Extra = 10
		} else {
			self.Extra = user.Domain.Extra + 1
		}

		self.Domain = self.Base + strconv.Itoa(self.Extra)
		return true
	}

	return true
}

func (self *UserDomain) Validate() bool {
	if self.Base == "" {
		self.AddError("domain.base cannot be empty")
		return false
	}

	if self.Extra == 0 {
		self.AddError("domain.extra cannot be empty")
		return false
	}

	if self.Extra < 10 {
		self.AddError("domain.extra cannot be less than 10")
		return false
	}

	if self.Domain == "" {
		self.AddError("domain.domain cannot be empty")
		return false
	}

	return true
}
