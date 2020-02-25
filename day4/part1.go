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

func readValidPasswords(readChannel, writeChannel chan int) {
	var count int
	for range readChannel {
		count++
	}
	writeChannel <- count
}

func findNumValidPasswordsInRange(lower, upper int) (numValidPasswords int) {
	countChannel := make(chan int)

	const n_threads = 1

	for i := 0; i < n_threads; i++ {
		// TODO: In order for n_threads to be > 1 we would need to have a way to divide the interval [lower, upper]
		tempChannel := make(chan int)
		go writeValidPasswordsInRange(tempChannel, lower, upper)
		go readValidPasswords(tempChannel, countChannel)
	}

	for i := 0; i < n_threads; i++ {
		numValidPasswords += <-countChannel
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
