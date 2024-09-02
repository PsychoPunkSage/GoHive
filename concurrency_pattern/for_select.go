package main

import (
	"fmt"
	"time"
)

// func main() {
// 	for_select()
// 	// infGoRoutine()
// }

func for_select() {
	forChannel := make(chan string, 7) // buffered channel
	chars := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, char := range chars {
		// select {
		// case forChannel <- char:
		// }
		forChannel <- char
	}

	close(forChannel)

	for ch := range forChannel {
		fmt.Println(ch)
	}
}

func infGoRoutine() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Hello")
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
