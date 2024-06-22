package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		// Gets `mutable` reference to `sum``
		sum += i
		return sum
	}
}

func closures() {
	countadder := adder()
	fmt.Println("+1", countadder(1))
	fmt.Println("+10", countadder(10))
	fmt.Println("+100", countadder(100))
}
