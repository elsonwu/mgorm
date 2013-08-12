package main

import (
	// "fmt"
	"github.com/elsonwu/mgorm"
	// "reflect"
)

type User struct {
	mgorm.Model `bson:",inline" json:",inline"`
	FullName    string      `bson:"fullname" json:"fullname"`
	FirstName   string      `bson:"first_name" json:"first_name"`
	Password    string      `bson:"password" json:"-"`
	LastName    string      `bson:"last_name" json:"last_name"`
	Email       string      `bson:"email" json:"email"`
	Profile     UserProfile `bson:"profile" json:"profile"`
}

func (self *User) Init() {
	self.Model.Init()
	self.Model.SetObj(self)
	self.Model.SetCollectionName(self.CollectionName())
}

func (self *User) CollectionName() string {
	return "user"
}
