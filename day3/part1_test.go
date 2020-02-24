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
			return
		}
	}
}

func TestPath2Vec(t *testing.T) {
	path := []string{"U2", "D5", "L1", "R6"}
	expected := [][2]int{
		{0, 2},
		{0, -5},
		{-1, 0},
		{6, 0},
	}
	actual := path2Vec(path)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d on input %s, got %d instead", expected, path, actual)
		return
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Expected %d on input %s, got %d instead", expected, path, actual)
			return
		}
	}
}
