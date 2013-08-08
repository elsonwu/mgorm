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
	self.Model.SetCollectionName(self.CollectionName())
}

func (self *User) CollectionName() string {
	return "user"
}

func (self *User) Validate() bool {
	if !self.Model.Validate() {
		return false
	}

	if !self.Profile.Validate() {
		return false
	}

	// refType := reflect.TypeOf(self)
	// refValue := reflect.ValueOf(self)
	// if refType.Kind() == reflect.Ptr {
	// 	refType = refType.Elem()
	// 	refValue = refValue.Elem()
	// }

	// numField := refType.NumField()
	// for i := 0; i < numField; i++ {
	// 	field := refType.Field(i)
	// 	if reflect.Struct != field.Type.Kind() || "Model" == field.Name {
	// 		continue
	// 	}

	// 	if v, ok := refValue.Field(i).Interface().(mgorm.IValidater); ok {
	// 		if !v.Validate() {
	// 			return false
	// 		}
	// 	}
	// }

	return true
}
