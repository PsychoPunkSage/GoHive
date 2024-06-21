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

func advanced_func() {
	fmt.Println("add:", proxy(2, 3, 4, add))
	fmt.Println("add:", proxy(2, 3, 4, mul))
}
