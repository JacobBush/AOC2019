package main

import "fmt"

func solveNounVerb(program []int, desiredOutput int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			progcopy := make([]int, len(program))
			copy(progcopy, program)
			progcopy = replaceInitialState(progcopy, noun, verb)
			progcopy = solveIntCode(progcopy, 0)
			if progcopy != nil && progcopy[0] == desiredOutput {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func part2() {
	program := readInput()
	noun, verb := solveNounVerb(program, 19690720)
	fmt.Printf("100 * noun + verb = %d\n", 100*noun+verb)
}
