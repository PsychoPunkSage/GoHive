package main

import (
	"fmt"
)

func main() {
	mychannel := make(chan string)

	go func() {
		mychannel <- someFunc("12")
	}()

	msg := <-mychannel

	fmt.Println("Message:> ", msg)

	fmt.Println("Hi")
}

func someFunc(num string) string {
	return num + "test"
}
