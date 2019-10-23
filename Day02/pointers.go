package main

import "fmt"

//Pointers in arithmetic

func change(numbers []int) {
	numbers[0] = 10000
}

//NOT RECOMMENDED
func changeUsingPointer(numbers *[]int) {
	(*numbers)[0] = 30000
	//numbers++
}

func main() {
	arr := []int{ 1, 2, 3 }
	changeUsingPointer(&arr)
	fmt.Println(arr[0])
	change(arr[:])
	change(arr)
	fmt.Println(arr[0])

	var p1 *int = new(int)
	fmt.Println(p1, *p1)
	*p1 = 10
	fmt.Println(p1, *p1)

	var pointerToFloat *float64 = new(float64)
	fmt.Println(pointerToFloat, *pointerToFloat)

	var pointer *string = new(string)
	fmt.Println(pointer, *pointer)

	x := 10
	fmt.Println(x, &x)
	//ptr is a pointer to x
	var ptr *int = &x
	fmt.Println(ptr, *ptr)
	incByRef(ptr)
	fmt.Println(x)

	y := 100
	incByRef(&y)
	fmt.Println(y)

	str := "cool"
	var p *string = &str
	//p++
	fmt.Println(p, *p)
	//inc(x)
	//fmt.Println(x)
}
func incByRef(number *int) {
	*number++
}



func inc(number int) {
	number++
}
