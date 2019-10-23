package main

import "fmt"

type Cat struct {
	name string
	age int
	cuteness int
}

type Nameless struct {
	string
	int64
}

type Author struct {
	name string
}

type Book struct {
	title string
	price float32
	Author
} 

func main() {
	playWithStructs()
	anonymous()
	nestedStruct()
}

func playWithStructs() {
	cat := Cat{
		name:     "Billy",
		age:      1,
		cuteness: 10,
	}
	fmt.Println(cat)
	fmt.Println(cat.name, cat.age, cat.cuteness)

	nameless := Nameless{string: "Patna", int64:  100000}
	fmt.Println(nameless)
	// takes default values
	nameless = Nameless{
		string: "Bangalore",
	}
	fmt.Println(nameless)
}


func anonymous() {
	anon := struct {
		name string
		age int
		address string
		ugliness int
	}{
		name:     "Anon",
		age:      33,
		address:  "Mother's basement",
	}
	fmt.Println(anon)
}

func nestedStruct() {
	book := Book{
		title:  "Seven Languages in Seven Days",
		price:  999.99,
		Author: Author{
			name: "Bruce Tate",
		},
	}
	fmt.Println(book)
}
