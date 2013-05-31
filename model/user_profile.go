package model

import (
	// "labix.org/v2/mgo"
	"api/model/attr"
)

type UserProfile struct {
	PrimaryEmail   attr.Email `bson:"primary_email" json:"primary_email"`
	SecondaryEmail attr.Email `bson:"secondary_email" json:"secondary_email"`
	Schools        []School   `bson:"schools" json:"schools"`
}
