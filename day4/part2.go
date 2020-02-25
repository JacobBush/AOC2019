package main

import (
	"fmt"
)

func part2() {
	PART = 2
	fmt.Println("Part 2")
	validPasswordsCount := findNumValidPasswordsInRange(lower, upper)
	fmt.Printf("There are %d valid passwords\n", validPasswordsCount)
}
