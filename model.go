package mgorm

import (
	// "fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Map map[string]interface{}

type IValidater interface {
	Validate() bool
}

type IEvent interface {
	On(event string, fn func(IModel) error)
	Emit(event string) error
}

type IModel interface {
	IErrorHandle
	IValidater
	IEvent
	IsNew() bool
	Init()
	GetId() bson.ObjectId
	AfterFind()
	BeforeSave() bool
	AfterSave()
	Collection() *mgo.Collection
	CollectionName() string
	SetCollectionName(name string)
	DB() *mgo.Database
	HasInited() bool
}

type Model struct {
	ErrorHandle    `bson:",inline" json:",inline"`
	Id             bson.ObjectId `bson:"_id" json:"id"`
	isNew          bool
	collectionName string
	inited         bool
	events         map[string][]func(IModel) error
}

func (self *Model) AfterFind() {
	self.isNew = false
}

func (self *Model) Validate() bool {
	self.ClearErrors()
	return true
}

func (self *Model) On(event string, fn func(IModel) error) {
	if nil == self.events {
		self.events = make(map[string][]func(IModel) error)
	}

	self.events[event] = append(self.events[event], fn)
}

func (self *Model) Emit(event string) error {
	if nil != self.events {
		length := len(self.events[event])
		for i := 0; i < length; i++ {
			err := self.events[event][i](self)
			if nil != err {
				return err
			}
		}
	}

	return nil
}

func (self *Model) BeforeSave() bool {
	return true
}

func (self *Model) AfterSave() {
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
