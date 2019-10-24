package playwithinterfaces

import "fmt"

type Person struct {
	name string
}
type Book struct {
	title string
}
type Reader interface {
	read() string
}
func doSomething(reader Reader) {
	fmt.Println(reader.read())
}
func (person Person) read() string {
	return person.name + " is reading a book"
}

func main() {
	p1 := Person{ "Sam" }
	fmt.Println(p1.read())
	var reader Reader = p1
	reader.read()
	doSomething(p1)

	//var reader2 Reader = Book { title: "ABC" }

}

