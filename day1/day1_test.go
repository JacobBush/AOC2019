package main

import "testing"

func TestFuelFromMass(t *testing.T) {
	masses := []int{12, 14, 1969, 100756}
	expected_fuels := []int{2, 2, 654, 33583}
	for i, mass := range masses {
		fuel := fuelFromMass(mass)
		if fuel != expected_fuels[i] {
			t.Errorf("Mass %d is correctly handled by fuelFromMass", mass)
		} else {
			t.Logf("Mass %d is correctly handled by fuelFromMass", mass)
		}
	}
}

func TestRecursiveFuel(t *testing.T) {
	masses := []int{14, 1969, 100756}
	expected_fuels := []int{2, 966, 50346}
	for i, mass := range masses {
		fuel := recursiveFuel(mass)
		if fuel != expected_fuels[i] {
			t.Errorf("Mass %d is incorrectly handled by recursiveFuel", mass)
		} else {
			t.Logf("Mass %d is correctly handled by recursiveFuel", mass)
		}
	}
}
