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

func (self *EmbeddedModel) Init() {

}

type Model struct {
	EmbeddedModel  `bson:",inline" json:"-"`
	Id             bson.ObjectId `bson:"_id" json:"id"`
	isOld          bool
	collectionName string
	inited         bool
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
	return !self.isOld
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
