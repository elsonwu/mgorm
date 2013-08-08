package mgorm

import (
	// "fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type IModel interface {
	IErrorHandle
	IValidater
	IEvent
	IsNew() bool
	Init()
	GetId() bson.ObjectId
	AfterFind()
	BeforeSave() error
	AfterSave()
	Collection() *mgo.Collection
	CollectionName() string
	SetCollectionName(name string)
	DB() *mgo.Database
	HasInited() bool
}

type Model struct {
	IValidater
	ErrorHandle    `bson:",inline" json:",inline"`
	Event          `bson:",inline" json:",inline"`
	Id             bson.ObjectId `bson:"_id" json:"id"`
	isNew          bool
	collectionName string
	inited         bool
}

func (self *Model) AfterFind() {
	self.Emit("AfterFind")
	self.isNew = false
}

func (self *Model) Validate() bool {
	self.ClearErrors()
	return true
}

func (self *Model) BeforeSave() error {
	err := self.Emit("BeforeSave")
	if nil != err {
		self.AddError(err.Error())
		return err
	}

	return nil
}

func (self *Model) AfterSave() {
	self.Emit("AfterSave")
	self.isNew = false
}

func (self *Model) Init() {
	self.inited = true
}

func (self *Model) HasInited() bool {
	return self.inited
}

func (self *Model) GetId() bson.ObjectId {
	return self.Id
}

func (self *Model) IsNew() bool {
	return self.isNew
}

func (self *Model) SetCollectionName(name string) {
	self.collectionName = name
}

func (self *Model) CollectionName() string {
	return self.collectionName
}

func (self *Model) DB() *mgo.Database {
	return DB()
}

func (self *Model) Collection() *mgo.Collection {
	if "" == self.CollectionName() {
		panic("collection name cannot be blank")
		return nil
	}

	return self.DB().C(self.CollectionName())
}
