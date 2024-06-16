package main

import "fmt"

func main() {
	var integer int8 = 12
	fmt.Printf("type of %d is %T\n", integer, integer)

	const pp_name = "PP" // Cannot alter `const`
	const bp_name = "BP"
	fmt.Printf("Premium Plan:  %s || Basic Plan:  %s\n", pp_name, bp_name)

	// Printf
	fmt.Printf("I am %v years old\n", 10)
	fmt.Printf("I am %v years old\n", "way too many")

	fmt.Printf("I am %s years old\n", "way too many")

	fmt.Printf("I am %d years old\n", 10)

	fmt.Printf("I am %.2f years old\n", 10.12)

	msg := fmt.Sprintf("this is an example of value being returned to a variable via `Sprintf` like::>%v", 100)
	fmt.Println(msg)

	// Concise if

	if len := getLength(12); len < 1 {
		fmt.Println("Invalid")
	} else {
		fmt.Printf("Valid")
	}
}

func getLength(rand int) int {
	return (rand/2 + 20)
}
