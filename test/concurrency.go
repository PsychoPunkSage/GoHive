package main

import (
	"fmt"
	"time"
)

/*
	time.Tick() : std lib func that returns a channel that sends a value on a given interval
	time.After() : sends a value once after the duration has passed.
	time.Sleep() : blocks a current goroutine for a specified amount of time.
*/

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Println("Email Received:", message)
	}()
	fmt.Println("Email Sent:", message)
}

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmails(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmails(email string) {
	fmt.Println("Email:", email)
}

func concurrenct() {
	msg_list := []string{"Yo AP is here", "Hello Everyone...", "incomming... runaway"}
	for _, msg := range msg_list {
		sendEmail(msg)
	}
}
