package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

func shapeMethod(shape shape) {
	fmt.Println("Area:", shape.area())
	fmt.Println("Perimeter:", shape.perimeter())
}

type employee interface {
	getName() string
	getSalary() float64
}

type contractor struct {
	name         string
	hourlyPay    float64
	hoursPerYear float64
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() float64 {
	return c.hourlyPay * c.hoursPerYear
}

func employeeInfo(emp employee) {
	fmt.Println("Name:", emp.getName())
	fmt.Println("Salary: $", emp.getSalary())
}

func interfaces() {
	cont := contractor{
		name:         "Contractor",
		hourlyPay:    12.94,
		hoursPerYear: 1234,
	}

	employeeInfo(cont)
}
