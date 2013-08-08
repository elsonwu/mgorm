#Simple Model

    package main

    import (
		"github.com/elsonwu/mgorm"
	)

	type User struct {
		mgorm.Model `bson:",inline" json:",inline"` //embed all base methods
		Username    string      `bson:"username" json:"username"`
		Email       string      `bson:"email" json:"email" rules:"email"`
	}

	func (self *User) Init() {
		self.Model.Init()
		self.Model.SetCollectionName(self.CollectionName())
	}

	func (self *User) CollectionName() string {
		return "user"
	}


#Find one

    user := new(User)
    err := mgorm.FindById(user, "51ffc45fad51987c28276e55")
    if nil != err {
	    fmt.Println(err)
    }
    
#Find List

	criteria := mgorm.NewCriteria()
	criteria.SetLimit(3)
	criteria.SetOffset(10)
	criteria.SetSelect([]string{"email"})
	criteria.AddSort("username", mgorm.CriteriaSortDesc)
	iter := mgorm.FindAll(user, criteria).Iter()
	users := make([]User, 3)
	iter.All(&users)
	fmt.Println(users)
	
#Save One

    user.Email = "test@gmail.com"
	err = mgorm.Save(user)
	if nil != err {
		fmt.Println(err)
	}	