package mgorm

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Map map[string]interface{}

type IModel interface {
	IErrorHandle
	IsNew() bool
	Init()
	GetId() bson.ObjectId
	AfterFind()
	Validate() bool
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
}

func (self *Model) AfterFind() {
	self.isNew = false
}

func (self *Model) Validate() bool {
	self.ClearErrors()
	return true
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
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("connect error:", err)
	}

	return session.DB("testcn10")
}

func (self *Model) Collection() *mgo.Collection {
	if "" == self.CollectionName() {
		panic("collection name cannot be blank")
		return nil
	}

	return self.DB().C(self.CollectionName())
}
