package main

import "testing"

func TestSeg2Vec(t *testing.T) {
	segs := []string{"U2", "D5", "L1", "R6"}
	expecteds := [][2]int{
		{0, 2},
		{0, -5},
		{-1, 0},
		{6, 0},
	}
	for i, seg := range segs {
		expected := expecteds[i]
		actual := seg2Vec(seg)
		if actual != expected {
			t.Errorf("Expected %d on input %s, got %d instead", expected, seg, actual)
		}
	}
}
