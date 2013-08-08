package main

import (
	"encoding/json"
	// "fmt"
	"github.com/elsonwu/mgorm"
	// "errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		user := new(User)

		//Find one:
		// err := mgorm.FindById(user, "51ffc45fad51987c28276e55")
		// if nil != err {
		// 	fmt.Println(err)
		// }

		// user.Email = "elsonwu@126.com"
		// if !mgorm.Save(user) {
		// 	fmt.Println("Save errors:", user.ErrorHandler().GetErrors())
		// }

		//Find list:
		criteria := mgorm.NewCriteria()
		criteria.SetLimit(3)
		criteria.AddSort("domain.domain", mgorm.CriteriaSortDesc)
		criteria.AddSort("domain.extra", mgorm.CriteriaSortAsc)
		iter := mgorm.FindAll(user, criteria).Iter()
		users := make([]User, 3)

		i := 0
		for iter.Next(user) {
			users[i] = *user
			i = i + 1
		}

		// fmt.Println(query)

		//user.Email = "test@126.com"
		// err = mgorm.Save(user)
		// if nil != err {
		// 	fmt.Println(err)
		// }

		output, _ := json.Marshal(users)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(output))
	})

	http.ListenAndServe(":8888", nil)
}
