package main

import (
	"github.com/elsonwu/mgorm/model"
)

type User struct {
	model.Model `bson:",inline" json:",inline"`
	FullName    string      `bson:"fullname" json:"fullname"`
	FirstName   string      `bson:"first_name" json:"first_name"`
	Password    string      `bson:"password" json:"-"`
	LastName    string      `bson:"last_name" json:"last_name"`
	Email       string      `bson:"email" json:"email" rules:"email"`
	Profile     UserProfile `bson:"profile" json:"profile"`
	Domain      UserDomain  `bson:"domain" json:"domain"`
}

func (self *User) Init() {
	self.Model.Init()
	self.Model.SetCollectionName(self.CollectionName())
}

func (self *User) CollectionName() string {
	return "user"
}
