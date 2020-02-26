package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() []int {
	csvFile, _ := os.Open("input.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	var program []int
	for _, i := range line {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, num)
	}
	return program
}

// func replaceInitialState(program []int, a int, b int) []int {
// 	program[1] = a
// 	program[2] = b
// 	return program
// }

func solveIntCode(program []int, position int) []int {
	if position >= len(program) {
		return nil
	}
	instruction := program[position]
	opcode := instruction % 100
	parameterModes := instruction / 100
	switch opcode {
	case 99:
		return program
	case 1:
		var mode1, mode2, mode3 int
		if parameterModes > 0 {
			mode1 = parameterModes % 10
			parameterModes /= 10
		}
		if parameterModes > 0 {
			mode2 = parameterModes % 10
			parameterModes /= 10
		}
		if parameterModes > 0 {
			mode3 = parameterModes % 10
		}
		var p1, p2, p3 int
		if mode1 == 1 {
			p1 = position + 1
		} else {
			p1 = program[position+1]
		}
		if mode2 == 1 {
			p2 = position + 2
		} else {
			p2 = program[position+2]
		}
		if mode3 == 1 {
			p3 = position + 3
		} else {
			p3 = program[position+3]
		}
		program[p3] = program[p2] + program[p1]
		return solveIntCode(program, position+4)
	case 2:
		var mode1, mode2, mode3 int
		if parameterModes > 0 {
			mode1 = parameterModes % 10
			parameterModes /= 10
		}
		if parameterModes > 0 {
			mode2 = parameterModes % 10
			parameterModes /= 10
		}
		if parameterModes > 0 {
			mode3 = parameterModes % 10
		}
		var p1, p2, p3 int
		if mode1 == 1 {
			p1 = position + 1
		} else {
			p1 = program[position+1]
		}
		if mode2 == 1 {
			p2 = position + 2
		} else {
			p2 = program[position+2]
		}
		if mode3 == 1 {
			p3 = position + 3
		} else {
			p3 = program[position+3]
		}
		program[p3] = program[p2] * program[p1]
		return solveIntCode(program, position+4)
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		text = strings.TrimSuffix(text, "\n")
		num, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		program[program[position+1]] = num
		return solveIntCode(program, position+2)
	case 4:
		var mode int
		if parameterModes > 0 {
			mode = parameterModes % 10
		}
		if mode == 0 {
			fmt.Printf("%d\n", program[program[position+1]])
		} else {
			fmt.Printf("%d\n", program[position+1])
		}
		return solveIntCode(program, position+2)
	default:
		fmt.Printf("Unknown opcode encountered: %d\n", program[position])
		return nil
	}
}

func part1() {
	program := readInput()
	// program = replaceInitialState(program, 12, 2)
	program = solveIntCode(program, 0)
	// fmt.Printf("Position 0 after program halts is: %d\n", program[0])
}
