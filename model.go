package mgorm

import (
	// "fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "reflect"
)

type IEmbeddedModel interface {
	IErrorHandler
	IValidator
	IEvent
}

type IModel interface {
	IEmbeddedModel
	IsNew() bool
	GetId() bson.ObjectId
	AfterFind()
	BeforeSave() error
	AfterSave()
	CollectionName() string
}

type EmbeddedModel struct {
	ErrorHandler `bson:",inline" json:"-"`
	Event        `bson:",inline" json:"-"`
}

func (self *EmbeddedModel) Validate() bool {
	self.ClearErrors()

	err := self.Emit("BeforeValidate")
	if nil != err {
		self.AddError(err.Error())
		return false
	}

	return true
}

type Model struct {
	EmbeddedModel  `bson:",inline" json:"-"`
	Id             bson.ObjectId `bson:"_id" json:"id"`
	isOld          bool
	collectionName string
}

func (self *Model) AfterFind() {
	self.Emit("AfterFind")
	self.isOld = true
}

func (self *Model) BeforeSave() error {
	if self.IsNew() {
		self.Id = bson.NewObjectId()
	}

	return self.Emit("BeforeSave")
}

func (self *Model) AfterSave() {
	self.Emit("AfterSave")
	self.isOld = true
}

func (self *Model) GetId() bson.ObjectId {
	return self.Id
}

func (self *Model) IsNew() bool {
	return !self.isOld
}
