package main

import "testing"

func TestSolveNounVerb(t *testing.T) {
	programs := [][]int{
		{1, -1, -1, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		{1, -1, -1, 0, 99},
	}

	expected_nv := [][2]int{
		{9, 10},
		{0, 0},
	}

	outs := []int{3500, 2}

	for i := range programs {
		noun, verb := solveNounVerb(programs[i], outs[i])
		enoun, everb := expected_nv[i][0], expected_nv[i][1]
		if noun != enoun || verb != everb {
			t.Errorf("Expected (%d, %d), got (%d, %d)", enoun, everb, noun, verb)
		}
	}
}
