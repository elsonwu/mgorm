package model

import (
	"fmt"
	"github.com/elsonwu/restapi/model/attr"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
)

type IDocument interface {
	GetCollectionName() string
}

type Document struct {
	Doc            IDocument     `bson:"" json:"-"`
	collectionName string        `bson:",omitempty" json:"-"`
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

func (self *Document) GetCollection() *mgo.Collection {
	if self.collectionName == "" {
		panic("the collection name is empty")
	}

	return self.Database().GetCollection(self.collectionName)
}

func (self *Document) GetFieldMapValue() attr.Map {
	return attr.Map{}
}

func (self *Document) Save() error {
	mapVal := make(bson.M)
	typ := reflect.TypeOf(self.Doc)
	val := reflect.ValueOf(self.Doc)
	for i := 0; i < typ.Elem().NumField(); i++ {
		f := typ.Elem().Field(i)
		if "_id" != f.Tag.Get("bson") {

			fmt.Println(f.Tag.Get("bson"))
			if !f.Anonymous {
				mapVal[f.Tag.Get("bson")] = val.Elem().Field(i).Interface()
			} else {
				fmt.Println(f.Type.NumField())
			}
		}
	}

	return self.GetCollection().Update(bson.M{"_id": self.Id}, bson.M{"$set": mapVal})
}
