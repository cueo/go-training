package main

func main() {
	shortIf()
	shortLoop()
	switchFn()
}

func shortIf() {
	/*
		x := 10
		if x % 2 == 0 {
			println(x, "is even")
		} else {
			println("x is odd")
		}
	*/

	if x := 10; x%2 == 0 {
		println("x is even")
	}

	// println(x)  // x os undefined
}

func shortLoop() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func switchFn() {
	switch x := 11; {
		case x%2 == 0:
			println("x is even")
		case x%2 != 0:
			println("x is odd")
		default:
			println("i don't know")
	}
}
