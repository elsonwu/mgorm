package model

import (
	//"labix.org/v2/mgo"
	"github.com/elsonwu/restapi/model/attr"
	// "labix.org/v2/mgo/bson"
)

type School struct {
	Document `json:",inline"`
	Name     attr.String `bson:"name" json:"name"`
	Type     attr.String `bson:"type" json:"type"`
}
