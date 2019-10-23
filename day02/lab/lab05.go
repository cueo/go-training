package main

import "fmt"

type Numbers []int

func main() {
	n := Numbers{5, 2, 6, 9, 1}
	fmt.Println("min:", n.min())
	fmt.Println("max:", n.max())
	fmt.Println("sum:", n.sum())
}

func (n Numbers) min() (m int) {
	m = n[0]
	for _, v := range n {
		if v < m {
			m = v
		}
	}
	return
}

func (n Numbers) max() (m int) {
	m = n[0]
	for _, v := range n {
		if v > m {
			m = v
		}
	}
	return
}

func (n Numbers) sum() (s int) {
	for _, v := range n {
		s += v
	}
	return
}

