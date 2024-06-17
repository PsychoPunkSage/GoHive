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

// Multiple interface
type expence interface {
	cost() float64
}
type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	emailAddress string
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return 10.00 * float64(len(e.emailAddress))
	}
	return 5 * float64(len(e.emailAddress))
}

func (e email) print() {
	fmt.Println("Email:", e.emailAddress)
}

type msg struct {
	sender  string
	message string
}

func (m msg) cost() float64 {
	if len(m.sender) != 0 {
		return 10.00 * float64(len(m.message))
	}
	return 5 * float64(len(m.message))
}

func (m msg) print() {
	fmt.Println("Email:", m.sender)
}

func emailExtraction(p printer, exp expence) {
	p.print()
	fmt.Println("COST:", exp.cost())
}

// Type Assertions
func getExpanceReport(e expence) (string, float64) {
	switch val := e.(type) {
	case msg:
		return val.sender, val.cost()
	case email:
		return val.emailAddress, val.cost()
	default:
		return "", 0.00
	}
}

func interfaces() {
	r := rect{width: 12.5, height: 5.0}
	shapeMethod(r)

	cont := contractor{
		name:         "Contractor",
		hourlyPay:    12.94,
		hoursPerYear: 1234,
	}

	employeeInfo(cont)

	e_mail := email{
		isSubscribed: false,
		emailAddress: "ap@ap.com",
	}
	msgs := msg{
		sender:  "PPAP",
		message: "Many things are there to take care of",
	}

	emailExtraction(e_mail, e_mail)

	// Type Assertion //
	a, b := getExpanceReport(e_mail)
	fmt.Println(a, b)
	a, b = getExpanceReport(msgs)
	fmt.Println(a, b)
}
