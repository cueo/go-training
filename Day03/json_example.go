package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//"id":1,"first_name":"Erskine","last_name":"Bushell","gender":"Male","income":1.36,"zip":"75500-000"

type Employee struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	Income float64 `json:"income"`
	Zip string `json:"zip"`
}

func main() {
		fileData, _ := ioutil.ReadFile("jsonfile")
		data := []byte(fileData)
//		fileData := []byte(`
//	  	[{
//			"id":1,
//			"first_name":"Erskine",
//			"last_name":"Bushell",
//			"gender":"Male",
//			"income":1.36,
//			"zip":"75500-000"
//		},
//{
//			"id":1,
//			"first_name":"Erskine",
//			"last_name":"Bushell",
//			"gender":"Male",
//			"income":1.36,
//			"zip":"75500-000"
//		},{
//			"id":1,
//			"first_name":"Erskine",
//			"last_name":"Bushell",
//			"gender":"Male",
//			"income":1.36,
//			"zip":"75500-000"
//		}]
//	`)
		var emp []Employee
		err := json.Unmarshal(data, &emp)
		if err == nil {
			fmt.Println(emp)
		} else {
			fmt.Println("Error", err)
		}
}
