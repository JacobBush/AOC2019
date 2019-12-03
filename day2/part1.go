package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []int {
	csvFile, _ := os.Open("input.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, error := reader.Read()
	if error != nil {
		log.Fatal(error)
	}
	var program []int
	for _, i := range line {
		num, error := strconv.Atoi(i)
		if error != nil {
			log.Fatal(error)
		}
		program = append(program, num)
	}
	return program
}

func replaceInitialState(program []int, a int, b int) []int {
	program[1] = a
	program[2] = b
	return program
}

func solveIntCode(program []int, position int) []int {
	switch program[position] {
	case 99:
		return program
	case 1:
		p1 := program[position+1]
		p2 := program[position+2]
		p3 := program[position+3]
		program[p3] = program[p2] + program[p1]
		return solveIntCode(program, position+4)
	case 2:
		// Some code duplication but whatever
		p1 := program[position+1]
		p2 := program[position+2]
		p3 := program[position+3]
		program[p3] = program[p2] * program[p1]
		return solveIntCode(program, position+4)
	default:
		fmt.Printf("Unknown opcode encountered: %d", program[position])
		return nil
	}
}

func part1() {
	program := readInput()
	program = replaceInitialState(program, 12, 2)
	program = solveIntCode(program, 0)
	fmt.Printf("Position 0 after program halts is: %d\n", program[0])
}
