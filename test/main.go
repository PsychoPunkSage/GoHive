package main

import (
	"fmt"
)

// Struct Methods
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width + r.height
}

type authentication struct {
	username string
	password string
}

func (auth authentication) basicAuth() string {
	return fmt.Sprintf("Authentication: Basic %s:%s", auth.username, auth.password)
}

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
		fmt.Println("Valid")
	}

	// functions
	fmt.Println(concat("AP ", "is here!!!"))

	// Memory assignment
	sendSoFar := 123
	const toAdd = 23
	sendSoFar = incrementSends(sendSoFar, toAdd)
	fmt.Println("You have sent", sendSoFar, "messages so far")

	// Anonymous Struct
	myCar := struct {
		company string
		model   string
	}{
		company: "BMW",
		model:   "i8",
	}

	fmt.Printf("%v\n", myCar)

	/////////////////////
	// Embeded Structs //
	/////////////////////

	// 	Here we take all the fields from one struct and shove it into the other...
	type Car struct {
		company string
		model   string
	}

	type Truck struct {
		Car
		Size int
	}

	lanesTruck := Truck{
		Size: 10,
		Car: Car{
			company: "BMW",
			model:   "i8",
		},
	}

	fmt.Println("Car name:", lanesTruck.company, lanesTruck.model)

	// Function for Struct
	r := rect{width: 12.5, height: 5.0}
	fmt.Println("Area:", r.area())

	a := authentication{username: "AP", password: "ap is here"}
	fmt.Println(a.basicAuth())

	////////////////
	// Interfaces //
	////////////////
	interfaces()

	////////////
	// Errors //
	////////////
	errors()
}

func getLength(rand int) int {
	return (rand/2 + 20)
}

func concat(s1, s2 string) string {
	return s1 + s2
}

// func take_another_finc(func(string, string) string, s string) string {
// 	return func(s1, s2 string) string {
//         return s1 + s2
//     }(s, s)
// }

func incrementSends(sendSoFar, toAdd int) int {
	sendSoFar += toAdd
	// fmt.Println(sendSoFar)
	return sendSoFar
}
