package main

import "testing"

func TestReadInput(t *testing.T) {
	program := readInput()
	if len(program) == 0 {
		t.Errorf("program not read in")
	}
}

// func TestReplaceInitialState(t *testing.T) {
// 	input := []int{9, 9, 9, 9}
// 	modifiedInput := replaceInitialState(input, 55, 3)
// 	expected := []int{9, 55, 3, 9}
// 	for i := range expected {
// 		if expected[i] != modifiedInput[i] {
// 			t.Errorf("state not modified properly")
// 		}
// 	}
// }

// func TestSolveIntCode(t *testing.T) {
// 	input := [][]int{
// 		{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
// 		{1, 0, 0, 0, 99},
// 		{2, 3, 0, 3, 99},
// 		{2, 4, 4, 5, 99, 0},
// 		{1, 1, 1, 4, 99, 5, 6, 0, 99},
// 	}

// 	expected := [][]int{
// 		{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
// 		{2, 0, 0, 0, 99},
// 		{2, 3, 0, 6, 99},
// 		{2, 4, 4, 5, 99, 9801},
// 		{30, 1, 1, 4, 2, 5, 6, 0, 99},
// 	}

// 	for i := range expected {
// 		in := input[i]
// 		out := solveIntCode(in, 0)
// 		for j := range out {
// 			if out[j] != expected[i][j] {
// 				t.Errorf("Program not solved correctly")
// 			}
// 		}
// 	}
// }
