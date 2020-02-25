package main

import (
	"fmt"
	"strconv"
)

func containsAdj(number string) bool {
	for i := range number {
		if i == 0 {
			continue
		}
		if number[i] == number[i-1] {
			return true
		}
	}
	return false
}

func alwaysIncr(number string) bool {
	for i := range number {
		if i == 0 {
			continue
		}
		a, _ := strconv.Atoi(string(number[i]))
		b, _ := strconv.Atoi(string(number[i-1]))
		if a < b {
			return false
		}
	}
	return true
}

func isValidPassword(password int) bool {
	strpass := strconv.Itoa(password)
	return containsAdj(strpass) && alwaysIncr(strpass)
}

func part1() {
	fmt.Println("Part 1")
}
