package main

import (
	"fmt"
	"time"
)

func main() {
	doneChannel()
}

func doneChannel() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(5 * time.Second)
	close(done)
}

func doWork( /* readonly channel */ done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work...")
		}
	}
}
