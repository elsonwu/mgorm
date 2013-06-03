package model

import (
	"github.com/elsonwu/restapi/model/attr"
	// "labix.org/v2/mgo"
	// "fmt"
	"labix.org/v2/mgo/bson"
	// "reflect"
)

func UserModel() (user *User) {
	user = new(User)
	user.Init()
	return
}

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
	if self.SetCollectionName(self.GetCollectionName()) {
		self.Document.doc = self
		self.UserDomain.owner = self
	}

	return self
}

func (self *User) New() (model *User) {
	model = new(User)
	model.Init()
	model.Document.isNew = true
	return
}

func (self *User) GetCollectionName() string {
	return "user"
}

func (self *User) BeforeSave() bool {
	if self.isNew {
		self.UserDomain.InitDomain()
	}

	if self.Email == "" {
		self.AddError("Email cannot be empty.")
		return false
	}

	if self.UserDomain.Base == "" {
		self.AddError("domain.base cannot be empty")
		return false
	}

	if self.UserDomain.Extra == 0 {
		self.AddError("domain.extra cannot be empty")
		return false
	}

	if self.UserDomain.Domain == "" {
		self.AddError("domain.domain cannot be empty")
		return false
	}

	return true
}

func (self *User) AfterFind() {
	self.Document.isNew = false
	self.DisplayName = self.FirstName + " " + self.LastName
}

func (self *User) FindAll(criteria Criteria) (models []*User, err error) {
	models = make([]*User, criteria.GetLimit())
	q := self.GetCollection().Find(criteria.GetConditions())
	q.Skip(criteria.GetOffset())
	q.Limit(criteria.GetLimit())
	q.All(&models)
	for k, _ := range models {
		models[k].Init()
		models[k].AfterFind()
	}
	return
}

func (self *User) Find(criteria Criteria) (*User, error) {
	q := self.GetCollection().Find(criteria.GetConditions())
	err := q.One(self)
	self.Init()
	self.AfterFind()
	return self, err
}

func (self *User) FindId(id string, criteria Criteria) (*User, error) {
	criteria.AddCondition("_id", "==", bson.ObjectIdHex(id))
	q := self.GetCollection().Find(criteria.GetConditions())
	err := q.One(self)
	self.Init()
	self.AfterFind()
	return self, err
}
