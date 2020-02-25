package main

import (
	"strconv"
	"testing"
)

func TestContainsAdj(t *testing.T) {
	input := []int{1111, 1234, 94423}
	expected := []bool{true, false, true}
	for i := range expected {
		actual := containsAdj(strconv.Itoa(input[i]))
		if expected[i] != actual {
			t.Errorf("Expected input %d, to produce output %t but was actually %t", input[i], expected[i], actual)
		}
	}
}

func TestAlwaysIncr(t *testing.T) {
	input := []int{1111, 1234, 94423}
	expected := []bool{true, true, false}
	for i := range expected {
		actual := alwaysIncr(strconv.Itoa(input[i]))
		if expected[i] != actual {
			t.Errorf("Expected input %d, to produce output %t but was actually %t", input[i], expected[i], actual)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	input := []int{111111, 223450, 123789}
	expected := []bool{true, false, false}
	for i := range expected {
		actual := isValidPassword(input[i])
		if expected[i] != actual {
			t.Errorf("Expected input %d, to produce output %t but was actually %t", input[i], expected[i], actual)
		}
	}
}
