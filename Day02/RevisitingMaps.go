package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4}
	fmt.Println(numbers)

	romanNumbers2  := map[int]string{1: "I", 2: "II"}
	fmt.Println(romanNumbers2)

	var romanNumbers map [int]string = make(map [int]string)
	romanNumbers[1] = "I"
	romanNumbers[2] = "II"
	romanNumbers[3] = "III"
	romanNumbers[4] = "IV"
	delete(romanNumbers, 3)

	for key, value := range romanNumbers   {
		fmt.Println(key, value)
	}

	//backtick operator
	multiLineString := `Hi!
			This is a paragraph!
		Cool! 
`
	fmt.Println(multiLineString)
}
