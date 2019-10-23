package main

import "fmt"

type Dog struct {
	name string
	age int
	floofness int
}

func main() {
	dog := Dog{}
	populateDog(dog)
	fmt.Println(dog)

	populateDogByRef(&dog)
	fmt.Println(dog)

	populateDogByRefDirect(&dog)
	fmt.Println(dog)
}

func populateDog(dog Dog) {
	dog.name = "Tucker"
	dog.age = 2
	dog.floofness = 10
}

func populateDogByRef(dog *Dog) {
	(*dog).name = "Kali"
	(*dog).age = 3
	(*dog).floofness = 10
}

func populateDogByRefDirect(dog *Dog) {
	dog.name = "Tommy"
	dog.age = 4
	dog.floofness = 10
}
