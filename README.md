> labix.org/v2/mgo -> gopkg.in/mgo.v2

> add `InitDBWithInfo` function

------

#Connect db
default

    mgorm.InitDB("127.0.0.1", "testcn10")
    
 with DialInfo
 
    dialInfo := &mgo.DialInfo{
        Addrs:    []string{"localhost:27017"},
        Timeout:  60 * time.Second,
        Database: "testcn10",
        Username: "usr",
        Password: "passwd1",
    }

    mgorm.InitDBWithInfo(dialInfo)

#Simple Model
All model are struct type

    package main

    import (
		"github.com/hangxin1940/mgorm"
	)

	type User struct {
		mgorm.Model `bson:",inline" json:",inline"` //embed all base methods
		Username    string      `bson:"username" json:"username"`
		Email       string      `bson:"email" json:"email"`
	}

	func (self *User) CollectionName() string {
		return "user"
	}

#Embedded Model

	type User struct {
		mgorm.Model `bson:",inline" json:",inline"` //embed all base methods
		Username    string      `bson:"username" json:"username"`
		Email       string      `bson:"email" json:"email"`
		Profile     UserProfile `bson:"profile" json:"profile"`
	}
	
    type UserProfile struct {
		mgorm.EmbeddedModel `bson:",inline" json:"-"` //use mgorm.EmbeddedModel
		SecondaryEmail      string `bson:"secondary_email" json:"secondary_email"`
		Website             string `bson:"website" json:"website"`
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
	
#Create One
    
    user := new(User)
    if !mgorm.Save(user) {
		fmt.Println(user.GetErrors())
	}
		
#Update One

    user.Email = "test@gmail.com"
	if !mgorm.Save(user) {
		fmt.Println(user.GetErrors())
	}	
	
#Error

	user.AddError("Test the error")
	if !mgorm.Save(user) {
		fmt.Println(user.GetErrors())
		//[Test the error]
	}
	
#Validate
	//use tag rules for field's validation
	type User struct {
		mgorm.Model `bson:",inline" json:",inline"` //embed all base methods
		Username    string      `bson:"username" json:"username"`
		Email       string      `bson:"email" json:"email" rules:"email"`
		Profile     UserProfile `bson:"profile" json:"profile"`
	}
	
    type UserProfile struct {
		mgorm.EmbeddedModel `bson:",inline" json:"-"` //use mgorm.EmbeddedModel
		SecondaryEmail      string `bson:"secondary_email" json:"secondary_email" rules:"email"`
		Website             string `bson:"website" json:"website" rules:"url"`
	}
	
	user := new(User)
	if mgorm.Validate(user) {
		fmt.Println(user.GetErrors(), user.Profile.GetErrors())
	}
	
	//When you run Save method, it will call Validate method automatically.
	if !user.Save() {
		fmt.Println(user.GetErrors(), user.Profile.GetErrors())
	}
	
	//You can also do more validate in your model, it will run when validating
	func (self *User) Validate() bool {
	
		//Don't forget to call the parent's Validate
	    if !self.Model.Validate() {
	    	return false
	    }
	    
	    if self.Username == "Admin" {
	    	self.AddError("You cannot use Admin as your username")
	    	return false
	    }
	    
	    return true
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
