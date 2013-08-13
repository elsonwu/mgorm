package main

import (
	// "fmt"
	// "errors"
	"github.com/elsonwu/mgorm"
	// "reflect"
)

type User struct {
	mgorm.Model `bson:",inline" json:",inline"`
	FullName    string       `bson:"fullname" json:"fullname"`
	FirstName   string       `bson:"first_name" json:"first_name"`
	Password    string       `bson:"password" json:"-"`
	LastName    string       `bson:"last_name" json:"last_name"`
	Email       string       `bson:"email" json:"email" rules:"email"`
	Profile     *UserProfile `bson:"profile" json:"profile"`
}

func (self *User) Init() {
	self.Model.Init()
	self.InitCollection()
	self.Profile.SetOwner(self)
}

func (self *User) Validate() bool {
	if !self.Model.Validate() {
		return false
	}

	return mgorm.NewValidator(self).Validate()
}

func (self *User) InitCollection() {
	self.Model.SetCollectionName(self.CollectionName())
}

func (self *User) CollectionName() string {
	return "user"
}
