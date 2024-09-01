package main

import (
	"fmt"
	"time"
)

func main() {
	go someFunc("12")
	go someFunc("13")
	go someFunc("14")

	time.Sleep(1 * time.Second)

	fmt.Println("Hi")
}

func someFunc(num string) {
	fmt.Println(num)
}
