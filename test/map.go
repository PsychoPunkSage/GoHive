package main

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	phoneNumber int
}

func fetUserMap(names []string, phoneNumber []int) (map[string]user, error) {
	usermap := make(map[string]user)
	if len(names) != len(phoneNumber) {
		return nil, errors.New("invalid Input")
	}
	for i := 0; i < len(names); i++ {
		usermap[names[i]] = user{name: names[i], phoneNumber: phoneNumber[i]}
	}

	return usermap, nil
}

func maps() {
	names := []string{"AP", "PPAP", "PPPP"}
	phoneNumber := []int{123456789, 987654321, 123456789}
	usermap, err := fetUserMap(names, phoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(usermap)
}
