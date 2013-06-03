package main

import (
	"encoding/json"
	"fmt"
	"github.com/elsonwu/restapi/model"
	// "github.com/elsonwu/restapi/model/attr"
	"net/http"
	// "reflect"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// data := model.UserModel().New()
		// data.FirstName = "elson"
		// data.LastName = "wu"
		// data.Email = "xx@111.com"
		// domain := *new(model.UserDomain)
		// domain.Base = "zzz"
		// domain.SetOwner(data)
		// data.UserDomain = domain
		// fmt.Println(data.Save())
		// fmt.Println(data.GetErrors())
		// fmt.Println(data.UserDomain)
		// //fmt.Printf("%+v", data)
		// return

		criteria := model.NewCriteria()
		data, _ := model.UserModel().FindId("51acaeb619d443a040cf0fd1", criteria)
		data.FirstName = "XXXXX"
		if !data.Save() {
			fmt.Println(data.GetErrors())
			res.Write([]byte("ERROR"))
		}
		output, err := json.Marshal(data)
		if err != nil {
			fmt.Println("marshal error")
			res.Write([]byte("json marshal error"))
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(output)
	})

	http.ListenAndServe(":8888", nil)
}
