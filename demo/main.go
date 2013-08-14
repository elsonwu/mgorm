package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"github.com/elsonwu/mgorm"
	"net/http"
)

func main() {
	mgorm.InitDB("127.0.0.1", "testcn10")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		user := new(User)
		user.Email = "xxx@xxx.com"
		fmt.Println(mgorm.Validate(user))
		fmt.Println(user.GetErrors(), user.Profile.GetErrors())

		// if !mgorm.Save(user) {
		// 	fmt.Println(user.GetErrors(), user.Profile.GetErrors())
		// }

		//Find one:
		// err := mgorm.FindById(user, "51ffc45fad51987c28276e55")
		// if nil != err {
		// 	fmt.Println(err)
		// }

		// user.FullName = "Admin"
		// user.Profile.PrimaryEmail = "tet@test.com"
		// user.Profile.Website = "test.com"

		// user.On("BeforeValidate", func() error {
		// 	if "Admin" == user.FullName {
		// 		return errors.New("You cannot use Admin")
		// 	}

		// 	return nil
		// })

		// user.FullName = "Admin"
		// user.On("TestEvent", func() error {
		// 	if "Admin" == user.FullName {
		// 		return errors.New("You cannot use Admin")
		// 	}
		// 	return nil
		// })

		// err = user.Emit("testevent")
		// if nil != err {
		// 	fmt.Println(err)
		// }

		// mgorm.Update(user, mgorm.Map{"email": "eee@eee.com"})

		// criteria := mgorm.NewCriteria()
		// criteria.AddCond("fullname", "==", "elson wu")
		// //mgorm.FindAll(user, criteria).One(user)
		// err = mgorm.Find(user, criteria)
		// fmt.Println(err)
		// fmt.Println(user)

		// user.Email = "elsonwu@126.com"
		// if !mgorm.Save(user) {
		// 	fmt.Println("Save errors:", user.ErrorHandler().GetErrors())
		// }

		//Find list:
		// criteria := mgorm.NewCriteria()
		// criteria.SetLimit(3)
		// criteria.AddSort("domain.domain", mgorm.CriteriaSortDesc)
		// criteria.AddSort("domain.extra", mgorm.CriteriaSortAsc)

		// users := make([]User, 3)
		// for i := 0; i < 3; i++ {
		// 	users[i].On("AfterFind", func() error {
		// 		fmt.Println("after find")

		// 		return nil
		// 	})
		// }

		// iter := mgorm.FindAll(user, criteria).Iter()
		// user.On("AfterFind", func() error {
		// 	fmt.Println("after find")
		// 	return nil
		// })

		// i := 0
		// for iter.Next(user) {
		// 	user.AfterFind()
		// 	users[i] = *user
		// 	i = i + 1
		// }

		//fmt.Println("validate:", users[0].Validate())
		//fmt.Println("errors:", users[0].GetErrors())

		//user.Email = "test@126.com"
		// err = mgorm.Save(user)
		// if nil != err {
		// 	fmt.Println(err)
		// }

		output, _ := json.Marshal(user)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(output))
	})

	http.ListenAndServe(":8888", nil)
}
