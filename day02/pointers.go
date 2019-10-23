package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x, &x, *&x)

	// noinspection GoVarAndConstTypeMayBeOmitted
	var addr *int = &x
	fmt.Printf("%T\n", addr)

	fmt.Println(x)
	inc(x)
	fmt.Println(x)
	incByRef(&x)
	fmt.Println(x)

	// addr++  // !! pointer arithmetic not allowed !!

	var strAddr *string
	fmt.Println(strAddr)  // <nil>
	// fmt.Println(*strAddr)  // invalid memory address or nil pointer dereference

	// noinspection GoVarAndConstTypeMayBeOmitted
	var strAddr2 *string = new(string)  // similar to malloc in C
	fmt.Println(strAddr2, *strAddr2, *strAddr2 == "")
}

func inc(a int) {
	a++
}

func incByRef(a *int) {
	*a++
}
