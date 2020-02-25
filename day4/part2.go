package main

import (
	"fmt"
)

func threeOrLonger(str string) bool {
	for i := range str {
		if i <= 1 {
			continue
		}
		if str[i] == str[i-1] && str[i] == str[i-2] {
			return true
		}
	}
	return false
}

func part2() {
	PART = 2
	fmt.Println("Part 2")
	validPasswordsCount := findNumValidPasswordsInRange(lower, upper)
	fmt.Printf("There are %d valid passwords\n", validPasswordsCount)
}
