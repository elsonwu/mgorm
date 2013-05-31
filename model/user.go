package model

import (
	"github.com/elsonwu/restapi/model/attr"
	// "labix.org/v2/mgo"
	// "fmt"
	"labix.org/v2/mgo/bson"
	// "reflect"
)

type User struct {
	Document    `bson:",inline" json:",inline"`
	FirstName   attr.String `bson:"first_name" json:"first_name"`
	Password    attr.String `bson:"password" json:"-"`
	LastName    attr.String `bson:"last_name" json:"last_name"`
	Email       attr.String `bson:"email" json:"email"`
	DisplayName attr.String `bson:"display_name" json:"display_name"`
	UserProfile UserProfile `bson:"profile" json:"profile"`
	UserDomain  UserDomain  `bson:"domain" json:"domain"`
}

func (self *User) Init() *User {
	if self.Document.collectionName == "" {
		self.Document.collectionName = self.GetCollectionName()
		self.Document.Doc = self
	}

	return self
}

func (self *User) Model() *User {
	return self.Init()
}

func (self *User) GetCollectionName() string {
	return "user"
}

func (self *User) AfterFind() {
	self.DisplayName = self.FirstName + " " + self.LastName
}

func (self *User) FindAll() (models []*User, err error) {
	models = make([]*User, 10)
	self.GetCollection().Find(bson.M{}).Limit(10).All(&models)
	return
}

func (self *User) New() (model *User) {
	model = new(User)
	model.Init()
	return
}

func (self *User) Find() (*User, error) {
	err := self.GetCollection().Find(bson.M{}).One(self)
	self.Init()
	self.AfterFind()
	return self, err
}

func (self *User) FindId(id string) (*User, error) {
	err := self.GetCollection().FindId(bson.ObjectIdHex(id)).One(self)
	self.Init()
	self.AfterFind()
	return self, err
}
