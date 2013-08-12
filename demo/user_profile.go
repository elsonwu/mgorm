package main

import (
// "fmt"
)

type UserProfile struct {
	PrimaryEmail   string `bson:"primary_email" json:"primary_email"`
	SecondaryEmail string `bson:"secondary_email" json:"secondary_email"`
}

func (self *UserProfile) Validate() bool {
	return true
}
