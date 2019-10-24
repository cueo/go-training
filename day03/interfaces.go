package main

import "fmt"

type Dog struct {
	name string
	age int
	floof int
}

type DogShow interface {
	showDog() string
}

func (dog Dog) showDog() string {
	return fmt.Sprintf("Dog called %s is %d years old and is %d/10 floofy.", dog.name, dog.age, dog.floof)
}

func main() {
	ptr := Dog{
		name:  "Kali",
		age:   3,
		floof: 10,
	}
	var dog DogShow
	dog = ptr
	fmt.Println(dog.showDog())
}
