package main

import "fmt"




type Person struct {
	name string
	age int
	city City
}
type City struct {
	name string
	population int64
}

type Data struct {
	string
	int64
}
type Book struct {
	title string
	price float64
	//author Author
	Author //Exported field
}
type Author struct {
	name string
}
func main() {
	b1 := Book{
		title: "ABC",
		price: 34.23,
		Author: Author{ name: "Mary"}}
	fmt.Println(b1.title, b1.price, b1.name)






	anony := struct{
		someData string
		someNumber int
	}{"value", 213123}
	fmt.Println(anony)


	d1 := Data{ "info", 324098}
	fmt.Println(d1, d1.string, d1.int64)

	p1 := Person{ name: "Sam", age: 12 }
	fmt.Println(p1.name, p1.age)

	p2 := Person{ "Ram", 23, City{"Chennai", 392840} }
	p3 := Person{}
	p4 := Person{age: 34, name: "Joe"}
	fmt.Println(p1, p2, p3, p4)
}






