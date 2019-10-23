package main

import "fmt"

func main()  {
	a := []int{9, 9, 6}
	fmt.Println(a)
	passByRef(a)
	fmt.Println(a)

	passByPointer(&a)
	fmt.Println(a)
}

func passByRef(a []int) {
	a[0] = 6
}

// ! not recommended !
func passByPointer(p *[]int) {
	(*p)[0] = 9
	// p++  // !! not allowed !!
}
