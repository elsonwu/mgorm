package model

import (
	"api/model/attr"
	// "labix.org/v2/mgo"
	// "fmt"
	"api/model/attr"
	// "labix.org/v2/mgo/bson"
	// "reflect"
)

type UserDomain struct {
	Base   attr.String `bson:"base" json:"base"`
	Extra  attr.Int    `bson:"extra" json:"extra"`
	Domain attr.String `bson:"domain" json:"domain"`
}
