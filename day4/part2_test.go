package main

import (
	"strconv"
	"testing"
)

func TestThreeOrLonger(t *testing.T) {
	PART = 2
	input := []int{112233, 123444, 111122}
	expected := []bool{true, false, true}
	for i := range expected {
		actual := threeOrLonger(strconv.Itoa(input[i]))
		if expected[i] != actual {
			t.Errorf("Expected input %d, to produce output %t but was actually %t", input[i], expected[i], actual)
		}
	}
}
