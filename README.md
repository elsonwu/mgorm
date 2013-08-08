#Simple Model
All model are struct type

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


#FindById

    user := new(User)
    err := mgorm.FindById(user, "51ffc45fad51987c28276e55")
    if nil != err {
	    fmt.Println(err)
    }
    
#Find One
	user := new(User)
    criteria := mgorm.NewCriteria()
    criteria.AddCond("username", "==", "elsonwu")
    mgorm.Find(user, criteria)
	fmt.Println(user)

#Find List

	criteria := mgorm.NewCriteria()
	criteria.SetLimit(3)
	criteria.SetOffset(10)
	criteria.SetSelect([]string{"email"})
	criteria.AddSort("username", mgorm.CriteriaSortDesc)
	users := make([]User, 3)
	mgorm.FindAll(user, criteria).All(&users)
	fmt.Println(users)
	
#Save One

    user.Email = "test@gmail.com"
	err = mgorm.Save(user)
	if nil != err {
		fmt.Println(err)
	}	
	
#Error
	user.AddError("Test the error")

	if !mgorm.Save(user) {
		fmt.Println(user.GetErrors())
		//[Test the error]
	}
	
#Event
##built-in event
	user.Username = "Admin"
	
	//The event name is case insensitive, so here you can also use "beforesave"
	user.On("BeforeSave", func() error {
		if "Admin" == user.Username {
			return errors.New("You cannot use Admin")
		}
		return nil
	})

	if !mgorm.Save(user) {
		fmt.Println(user.GetErrors())
		//[You cannot use Admin]
	}
	
	
##Customized event
	//You can emit the event manually
	
	user.Username = "Admin"
	user.On("TestEvent", func() error {
	    if "Admin" == user.Username {
			return errors.New("You cannot use Admin")
		}
		return nil
	})
	
	//Notice: the event name is case insensitive
	err := user.Emit("testevent")
	if nil != err {
		fmt.Println(err)
		//You cannot use Admin
	}