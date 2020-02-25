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

func writeValidPasswordsInRange(ch chan int, lower, upper int) {
	for i := lower; i <= upper; i++ {
		if isValidPassword(i) {
			ch <- i
		}
	}
	close(ch)
}

func findNumValidPasswordsInRange(lower, upper int) (numValidPasswords int) {
	ch := make(chan int)
	go writeValidPasswordsInRange(ch, lower, upper)
	for range ch {
		numValidPasswords++
	}
	return
}

func part1() {
	const lower = 372304
	const upper = 847060
	fmt.Println("Part 1")
	validPasswordsCount := findNumValidPasswordsInRange(lower, upper)
	fmt.Printf("There are %d valid passwords\n", validPasswordsCount)
}
