package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Cuteness int `json:"cuteness"`
}

func main() {
	s := []byte(`{
  "name": "Billy",
  "age": 1,
  "cuteness": 10
}`)

	var cat Cat
	err := json.Unmarshal(s, &cat)
	if err == nil {
		fmt.Println(cat)
	}
}
