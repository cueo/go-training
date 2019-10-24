package yesterday

import "fmt"

type Person struct {
	Name string
	Age int
}

//func PrintDetails(person Person){}
func (person* Person) PrintDetails() {
	fmt.Println(person.Name, person.Age)
}