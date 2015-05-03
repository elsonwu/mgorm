package main

import (
	// "fmt"
	// "errors"
	"github.com/hangxin1940/mgorm"
)

type UserProfile struct {
	mgorm.EmbeddedModel `bson:",inline" json:"-"`
	PrimaryEmail        string `bson:"primary_email" json:"primary_email" rules:"email"`
	SecondaryEmail      string `bson:"secondary_email" json:"secondary_email" rules:"email"`
	Website             string `bson:"website" json:"website" rules:"url"`
}
