package main

func main() {
	var x int = 10
	// curly braces are compulsory!
	if x % 2 == 0 {
		println(x, "is even")
	} else { // else has to be inline (because no ;)
		println("x is odd")
	}

	// only for keyword (no while, no do..while)
	var count = 0
	for count < x {
		println("Hello, world!")
		count++
		// ++count  // no pre-increment, decrement
	}

	count = 0
	for ; count < x; count++  {  // does not allow parantheses
		println("Hello, universe!")
	}
}
