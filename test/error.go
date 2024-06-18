package main

import (
	"errors"
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost_customer, err := sendSMS(msgToCustomer)
	if err != nil {
		fmt.Println(err)
		return 0.0, err
	}

	cost_Spouse, err := sendSMS(msgToSpouse)
	if err != nil {
		fmt.Println(err)
		return 0.0, err
	}

	return cost_customer + cost_Spouse, nil
}

func sendSMS(msg string) (float64, error) {
	const maxTextLen = 25
	const costPerCharacter = 0.02
	if len(msg) > maxTextLen {
		return 0.00, fmt.Errorf("message is too long")
	}
	return float64(len(msg)) * costPerCharacter, nil
}

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by 0", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.00, divideError{dividend}
	}

	if divisor == 1 {
		return dividend, errors.New("No dividing by 1")
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

// MAIN \\
func errors_package() {
	cost, err := sendSMSToCouple("hu0qwu9-0u10u123092", "8921yr9 h2ye02u091y012")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cost)

	// ---------------
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
