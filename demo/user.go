package main

import (
	// "fmt"
	// "errors"
	"github.com/hangxin1940/mgorm"
	// "reflect"
)

type User struct {
	mgorm.Model `bson:",inline"`
	FullName    string      `bson:"fullname" json:"fullname"`
	FirstName   string      `bson:"first_name" json:"first_name"`
	Password    string      `bson:"password" json:"-"`
	LastName    string      `bson:"last_name" json:"last_name"`
	Email       string      `bson:"email" json:"email" rules:"email"`
	Profile     UserProfile `bson:"profile" json:"profile"`
}

func (self *User) CollectionName() string {
	return "user"
}
