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
		criteria := model.NewCriteria()
		criteria.SetLimit(2)
		data, _ := model.UserModel().FindAll(criteria)
		// err := data.Save()
		// if err != nil {
		// 	fmt.Println(err)
		// 	res.Write([]byte("ERROR"))
		// }
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
