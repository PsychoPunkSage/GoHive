package main

import "fmt"

func main() {
	pipelines()
}

func pipelines() {
	nums := []int{2, 5, 4, 8, 7, 10}

	// stage 1
	dataChannel := sliceToChannels(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
		fmt.Println("Yo")
	}
}

func sliceToChannels(nums []int) <-chan int {
	ch := make(chan int /*, len(nums)*/) // unbuffered channel
	go func() {                          // async : i.e. golang will not wait for it to be over
		for _, i := range nums {
			// select {
			// default:
			// 	ch <- i
			// }
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func sq(dataChannel <-chan int) <-chan int {
	ch := make(chan int)
	go func() { // async : i.e. golang will not wait for it to be over
		for n := range dataChannel {
			ch <- n * n
		}
		close(ch)
	}()
	return ch
}
