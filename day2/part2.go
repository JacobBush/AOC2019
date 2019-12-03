package main

import "fmt"

func solveNounVerb(desiredOutput int) (int, int) {
	program := readInput()
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			progcopy := make([]int, len(program))
			copy(progcopy, program)
			progcopy = replaceInitialState(progcopy, noun, verb)
			progcopy = solveIntCode(progcopy, 0)
			if progcopy[0] == desiredOutput {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func part2() {
	noun, verb := solveNounVerb(19690720)
	fmt.Printf("100 * noun + verb = %d\n", 100*noun+verb)
}
