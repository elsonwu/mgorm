package model

import (
	// "fmt"
	"labix.org/v2/mgo"
)

type ICollection interface {
	GetCollection() *mgo.Collection
	GetCollectionName() string
}
