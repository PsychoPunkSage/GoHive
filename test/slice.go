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

func slices() {
	total := summ(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("total: ", total)

	// Spread operator
	nums := []float64{1, 2, 3, 4, 5, 6, 7}
	total1 := summ(nums...)
	fmt.Println("total1: ", total1)

	cost := []cost{{0, 4.0}, {1, 2.1}, {1, 3.1}, {5, 2.5}}
	fmt.Println("cost: ", getCostByDay(cost))

}
