package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readInput() []int {
	csvFile, _ := os.Open("day1_input.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var masses []int
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		mass, error := strconv.Atoi(line[0])
		if error != nil {
			log.Fatal(error)
		}
		masses = append(masses, mass)
	}
	return masses
}

func fuelFromMass(mass int) int {
	var fuel int = (mass / 3) - 2
	return fuel
}

func recursiveFuel(mass int) int {
	fuel_required := fuelFromMass(mass)
	if fuel_required < 0 {
		return 0
	}
	fuel_for_fuel := recursiveFuel(fuel_required)
	total_fuel := fuel_required + fuel_for_fuel
	return total_fuel
}

func main() {

	masses := readInput()
	var sum_of_fuel int = 0
	var recursive_sum_of_fuel int = 0
	for _, mass := range masses {
		fuel_required := fuelFromMass(mass)
		sum_of_fuel += fuel_required
		recursive_fuel := recursiveFuel(mass)
		recursive_sum_of_fuel += recursive_fuel
	}
	fmt.Println("Part 1 fuel required is", sum_of_fuel)
	fmt.Println("Part 2 fuel required is", recursive_sum_of_fuel)
}
