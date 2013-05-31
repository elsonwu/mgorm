package model

import (
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type School struct {
	Document `json:",inline"`
	Name     string `bson:"name" json:"name"`
	Type     string `bson:"type" json:"type"`
}

func (self *School) Model() *School {
	school := School{}
	school.Document = Document{collectionName: "school"}
	return &school
}

func (self *School) FindId(id string) (model *School, err error) {
	err = self.GetCollection().FindId(bson.ObjectIdHex(id)).One(&model)
	return
}
