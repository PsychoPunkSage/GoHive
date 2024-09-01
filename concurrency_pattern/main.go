package main

import (
	"fmt"
)

func main() {
	mychannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		mychannel <- someFunc("12")
	}()

	go func() {
		anotherChannel <- someFunc("34")
	}()

	// msg := <-mychannel

	// fmt.Println("Message:> ", msg)

	// fmt.Println("Hi")

	select { // its gonna block until it receives info from one of the channels.
	case msgMychannel := <-mychannel:
		fmt.Println(msgMychannel)
	case msgAnotherChannel := <-anotherChannel:
		fmt.Println(msgAnotherChannel)
	}
}

func someFunc(num string) string {
	return num + "test"
}
