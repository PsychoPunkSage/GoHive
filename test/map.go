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

type data struct {
	name                string
	number              int
	scheduledForDeleted bool
}

func deleteIfNecessary(info map[string]data, name string) (deleted bool, err error) {
	// As you can se we didn't return the map...
	// But still the map is Altered by the opertion inside the function.
	data, ok := info[name]
	if !ok {
		return false, errors.New("Data not found")
	}
	if data.scheduledForDeleted {
		delete(info, name)
		return true, nil
	} else {
		return false, nil
	}
}

/*
Much simpler to use `struct` directly as key instead of `nested maps`
*/

func getCounts(userIds []string) map[string]int {
	counts := make(map[string]int)
	for _, id := range userIds {
		count, ok := counts[id]
		if !ok {
			counts[id] = 0
		}
		counts[id] = count + 1
	}
	return counts
}

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstchar := rune(name[0])
		if _, ok := counts[firstchar]; !ok {
			counts[rune(name[0])] = make(map[string]int)
		}
		counts[rune(name[0])][name]++
	}
	return counts
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
	fmt.Println(usermap["AP"])
	delete(usermap, "PPPP")
	fmt.Println("PPP:", usermap["PPP"]) // if a key DNE then GO returns 0....
	// Check deleted key
	ele, ok := usermap["PPPP"]
	if ok {
		fmt.Println(ele)
	} else {
		fmt.Println("Key not found")
	}
	fmt.Println(usermap)

	// ---------------------------------
	uid := []string{"a", "b", "c", "d", "a", "e", "f", "g", "h", "c", "c"}
	fmt.Println(getCounts(uid))

	// ---------------------------------
	names = []string{"aa", "aps", "aa", "bd", "bdc"}
	fmt.Println(getNameCounts(names))
}
