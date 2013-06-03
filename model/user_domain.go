package model

import (
	"fmt"
	"github.com/elsonwu/restapi/model/attr"
	"strconv"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	// "reflect"
)

type UserDomain struct {
	owner  *User
	Base   attr.String `bson:"base" json:"base"`
	Extra  attr.Int    `bson:"extra" json:"extra"`
	Domain attr.String `bson:"domain" json:"domain"`
}

func (self *UserDomain) SetOwner(user *User) {
	self.owner = user
}

func (self *UserDomain) InitDomain() bool {
	if self.owner.isNew {
		u := UserModel().New()
		self.owner.GetCollection().Find(attr.Map{"domain.base": self.Base}).Sort("+domain.extra").One(&u)
		if u.UserDomain.Extra == 0 {
			self.Extra = 10
		} else {
			self.Extra = u.UserDomain.Extra + 1
		}

		if self.Base == "" {
			self.owner.AddError("domain.base cannot be empty.")
			return false
		}

		self.Domain = *new(attr.String)
		self.Domain.Set(self.Base.Get() + strconv.Itoa(self.Extra.Get()))
		fmt.Println(self)
		return true
	}

	return true
}
