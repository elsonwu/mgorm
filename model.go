package mgorm

import (
	// "fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	// "reflect"
)

type IEmbeddedModel interface {
	IErrorHandler
	IValidator
	IEvent
	Init()
	SetOwner(IEmbeddedModel)
	GetOwner() IEmbeddedModel
}

type IModel interface {
	IEmbeddedModel
	IsNew() bool
	GetId() bson.ObjectId
	AfterFind()
	BeforeSave() error
	AfterSave()
	Collection() *mgo.Collection
	CollectionName() string
	InitCollection()
	SetCollectionName(name string)
	DB() *mgo.Database
	HasInited() bool
}

type EmbeddedModel struct {
	ErrorHandler `bson:",inline" json:",inline"`
	Event        `bson:",inline" json:",inline"`
	owner        IEmbeddedModel `bson:",inline" json:",inline"`
}

func (self *EmbeddedModel) SetOwner(model IEmbeddedModel) {
	self.owner = model
}

func (self *EmbeddedModel) GetOwner() IEmbeddedModel {
	return self.owner
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
	EmbeddedModel
	Id             bson.ObjectId `bson:"_id" json:"id"`
	isNew          bool
	collectionName string
	inited         bool
}

func (self *Model) AfterFind() {
	self.Emit("AfterFind")
	self.isNew = false
}

func (self *Model) BeforeSave() error {
	return self.Emit("BeforeSave")
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
