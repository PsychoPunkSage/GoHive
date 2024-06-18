package main

import "fmt"

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

func errors() {
	cost, err := sendSMSToCouple("hu0qwu9-0u10u123092", "8921yr9 h2ye02u091y012")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cost)
}
