package main

import (
	yes "./yesterday"
	"fmt"
)


func main() {
	arr := []int { 100, 200 }
	//Not really recommended
	var numbers *[]int = &[]int { 2, 3 }

	//numbers++
	fmt.Println(arr[0], (*numbers)[0])

	 p1 := &yes.Person{ "Sam", 12 }
	fmt.Println(p1.Name, (*p1).Name)
	 //PrintDetails(p1)
	 p1.PrintDetails() // p1 -> printDetails()

	 //(&p1).PrintDetails()

}
