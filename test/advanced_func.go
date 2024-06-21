package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func mul(x, y float64) float64 {
	return x * y
}

func proxy(a, b, c float64, proxy func(float64, float64) float64) float64 {
	return proxy(proxy(a, b), c)
}

// Currying
func selfMath(mathfunc func(float64, float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return mathfunc(x, x)
	}
}

func logger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}

func advanced_func() {
	fmt.Println("add:", proxy(2, 3, 4, add))
	fmt.Println("add:", proxy(2, 3, 4, mul))

	// currying
	sqFunc := selfMath(mul)
	doubleFunc := selfMath(add)

	fmt.Println(sqFunc(10))
	fmt.Println(doubleFunc(10))
}
