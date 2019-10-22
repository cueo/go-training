package main

func main() {
	//x := 10
	//if x % 2 == 0 {
	//	println("x is even")
	//}

	if x :=10; x % 2 == 0 {
		println("x is even")
	}
	switch x :=10;  {
		case x % 2 == 0:
			println("x is even")
		case x % 2 	== 1:
			println("x is odd")
		default:
			println("huh")
	}

	for i := 0; i < 5 ; i++  {
		println(i)
	}

	//Undefined
	//println(x)
}
