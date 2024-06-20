package main

import "fmt"

// Variadic function
func summ(nums ...float64) float64 {
	num := 0.00
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}

	return num
}

type cost struct {
	day  int
	cost float64
}

func getCostByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		for costs[i].day >= len(costByDay) {
			costByDay = append(costByDay, 0.00)
		}
		costByDay[costs[i].day] += costs[i].cost
	}
	return costByDay
}

// 2d slices
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func potential_error() {
	fmt.Println("================NO Error================")
	a := make([]int, 3)
	fmt.Println("Len(a) = ", len(a))
	fmt.Println("Cap(a) = ", cap(a))
	fmt.Println("Addr(a) = ", &a[0])

	fmt.Println("Appending 4 to b from a")
	b := append(a, 4)
	fmt.Println("b = ", b)
	fmt.Println("Addr(b) = ", &b[0])

	fmt.Println("Appending 5 to c from a")
	c := append(a, 5)
	fmt.Println("c = ", c)
	fmt.Println("Addr(c) = ", &c[0])

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	fmt.Println("================Error================")

	i := make([]int, 3, 5)
	fmt.Println("Len(i) = ", len(i))
	fmt.Println("Cap(i) = ", cap(i))
	fmt.Println("Addr(i) = ", &a[0])

	fmt.Println("Appending 4 to j from i")
	j := append(i, 4, 5, 6)
	fmt.Println("j = ", j)
	fmt.Println("Addr(j) = ", &j[0])

	fmt.Println("Appending 5 to k from i")
	k := append(i, 15, 16, 17)
	fmt.Println("k = ", k)
	fmt.Println("Addr(c) = ", &k[0])

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)

	fmt.Println("Always assign the value of `append()` fn back to the same slice")
}

func slices() {
	total := summ(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("total: ", total)

	// Spread operator
	nums := []float64{1, 2, 3, 4, 5, 6, 7}
	total1 := summ(nums...)
	fmt.Println("total1: ", total1)

	cost := []cost{{0, 4.0}, {1, 2.1}, {1, 3.1}, {5, 2.5}}
	fmt.Println("cost: ", getCostByDay(cost))

	fmt.Println("Matrix:\n", createMatrix(4, 4))

	potential_error()
}
