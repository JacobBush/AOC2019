package main

import (
	"fmt"
	"strconv"
)

// PART designates which part we are in
var PART int

func containsAdj(number string) bool {
	for i := range number {
		if i == 0 {
			continue
		}
		if number[i] == number[i-1] {
			if PART == 1 {
				return true
			}
			if i > 1 && number[i] == number[i-2] {
				continue
			}
			if i < (len(number)-1) && number[i] == number[i+1] {
				continue
			}
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

	const nThreads = 4
	stepSize := (upper - lower) / nThreads

	for i := 0; i < nThreads; i++ {
		tempChannel := make(chan int)
		sliceLower := lower + i*stepSize
		sliceUpper := lower + (i+1)*stepSize
		if i == nThreads-1 {
			// Correct for floor operation - last thread will take up to (nThreads - 1) extra
			sliceUpper = upper
		}
		go writeValidPasswordsInRange(tempChannel, sliceLower, sliceUpper)
		go readValidPasswords(tempChannel, countChannel)
	}

	for i := 0; i < nThreads; i++ {
		numValidPasswords += <-countChannel
	}

	return
}

const lower = 372304
const upper = 847060

func part1() {
	PART = 1
	fmt.Println("Part 1")
	validPasswordsCount := findNumValidPasswordsInRange(lower, upper)
	fmt.Printf("There are %d valid passwords\n", validPasswordsCount)
}
