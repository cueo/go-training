package structfunc

import "fmt"

type Cat struct {
	Name     string
	Age      int
	Cuteness int
}

func (cat Cat) ShowCat() {
	fmt.Println(cat.Name, cat.Age, cat.Cuteness)
}
