package model

import (
	"errors"
	//"fmt"
	"github.com/elsonwu/restapi/model/attr"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ISubDocument interface {
	Validate() bool
}

type IDocument interface {
	ISubDocument
	GetCollectionName() string
	BeforeSave() bool
	AfterSave()
}

type Document struct {
	ErrorHandle    `bson:",inline" json:"-"`
	doc            IDocument     `bson:",omitempty" json:"-"`
	collectionName string        `bson:",omitempty" json:"-"`
	isNew          bool          `bson:",omitempty" json:"-"`
	Id             bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Ctime          attr.Int      `bson:"ctime" json:"ctime"`
	Mtime          attr.Int      `bson:"mtime" json:"mtime"`
}

func (self *Document) IdString() string {
	return self.Id.Hex()
}

func (self *Document) Database() *Database {
	return GetDatabase()
}

func (self *Document) SetCollectionName(name string) bool {
	if self.collectionName == "" {
		self.collectionName = name
		return true
	}

	return false
}

func (self *Document) GetCollectionName() string {
	panic("please overrid this method in sub document")
	return "please overrid this method in sub document"
}

func (self *Document) Validate() bool {
	return true
}

func (self *Document) BeforeSave() bool {
	return true
}

func (self *Document) AfterSave() {}

func (self *Document) GetCollection() *mgo.Collection {
	if self.collectionName == "" {
		panic("the collection name is empty")
	}

	return self.Database().GetCollection(self.collectionName)
}

func (self *Document) GetFieldMapValue() attr.Map {
	return attr.Map{}
}

func (self *Document) Save() bool {

	if !self.doc.Validate() || !self.doc.BeforeSave() {
		return false
	}

	err := errors.New("")
	if self.isNew {
		err = self.GetCollection().Insert(self.doc)
	} else {
		err = self.GetCollection().Update(bson.M{"_id": self.Id}, self.doc)
	}

	if err != nil {
		if e, ok := err.(error); ok && e.Error() != "" {
			self.AddError(e.Error())
		}

		return false
	}

	self.doc.AfterSave()
	return true
}
