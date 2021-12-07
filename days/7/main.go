package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabs := loadCrabs()
	part1(crabs)
	part2(crabs)
}

func part1(crabs []int) int {
	min, max := findMaximaPositions(crabs)

	lowestFuelNeeded := math.MaxInt
	for position := min; position < max; position++ {

		var fuelNeeded int
		for _, crab := range crabs {
			switch {
			case position < crab:
				fuelNeeded += crab - position
			case position > crab:
				fuelNeeded += position - crab
			}
		}

		if fuelNeeded < lowestFuelNeeded {
			lowestFuelNeeded = fuelNeeded
		}
	}

	fmt.Printf("lowest fuel needed %d\n", lowestFuelNeeded)

	return lowestFuelNeeded
}

func part2(crabs []int) int {
	fuelRequired := func(currentPosition, targetPosition int) int {
		var steps int

		switch {
		case currentPosition > targetPosition:
			steps = currentPosition - targetPosition
		case currentPosition < targetPosition:
			steps = targetPosition - currentPosition
		}

		// https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_⋯
		return (steps * (steps + 1)) / 2
	}

	min, max := findMaximaPositions(crabs)

	lowestFuelNeeded := math.MaxInt
	for position := min; position < max; position++ {

		var fuelNeeded int
		for _, crab := range crabs {
			fuelNeeded += fuelRequired(crab, position)
		}

		if fuelNeeded < lowestFuelNeeded {
			lowestFuelNeeded = fuelNeeded
		}
	}

	fmt.Printf("lowest fuel needed %d\n", lowestFuelNeeded)

	return lowestFuelNeeded
}

func loadCrabs() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(f *os.File) {
		cerr := f.Close()
		if cerr != nil {
			fmt.Printf("file: %v\n", cerr)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	numbers := strings.Split(line, ",")

	var crabs []int
	for _, number := range numbers {
		n, perr := strconv.ParseInt(number, 10, 16)
		if perr != nil {
			panic(perr)
		}
		crabs = append(crabs, int(n))
	}

	return crabs
}

func findMaximaPositions(crabs []int) (int, int) {
	min := math.MaxInt
	max := 0

	for _, crab := range crabs {
		switch {
		case crab > max:
			max = crab
		case crab < min:
			min = crab
		}
	}

	return min, max
}
