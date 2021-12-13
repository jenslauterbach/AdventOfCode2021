package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	octopuses := loadOctopuses()
	part1(octopuses)
	round := part2(octopuses)
	fmt.Printf("round: %d\n", round)
}

func part1(octopuses [12][12]int) int {
	var flashes int

	for round := 1; round <= 100; round++ {

		// Step 1: Increase energy of every octopus by 1
		for y := 1; y < len(octopuses)-1; y++ {
			for x := 1; x < len(octopuses[y])-1; x++ {
				octopuses[y][x]++
			}
		}

		// Step 2: Process glowing octopus
		for {
			if !hasOctopusWithHighEnergy(octopuses) {
				break
			}

			for y := 1; y < len(octopuses)-1; y++ {
				for x := 1; x < len(octopuses[y])-1; x++ {
					if octopuses[y][x] < 10 {
						continue
					}

					// Increase glow count:
					flashes++
					// Set to 0:
					octopuses[y][x] = 0

					if octopuses[y-1][x-1] > 0 {
						octopuses[y-1][x-1]++
					}
					if octopuses[y-1][x] > 0 {
						octopuses[y-1][x]++
					}
					if octopuses[y-1][x+1] > 0 {
						octopuses[y-1][x+1]++
					}
					if octopuses[y][x-1] > 0 {
						octopuses[y][x-1]++
					}
					if octopuses[y][x+1] > 0 {
						octopuses[y][x+1]++
					}
					if octopuses[y+1][x-1] > 0 {
						octopuses[y+1][x-1]++
					}
					if octopuses[y+1][x] > 0 {
						octopuses[y+1][x]++
					}
					if octopuses[y+1][x+1] > 0 {
						octopuses[y+1][x+1]++
					}
				}
			}
		}
	}

	fmt.Printf("flashes: %d\n", flashes)
	return flashes
}

func part2(octopuses [12][12]int) int {
	for round := 1; ; round++ {
		if allOctopusesFlashed(octopuses) {
			return round - 1
		}

		// Step 1: Increase energy of every octopus by 1
		for y := 1; y < len(octopuses)-1; y++ {
			for x := 1; x < len(octopuses[y])-1; x++ {
				octopuses[y][x]++
			}
		}

		// Step 2: Process glowing octopus
		for {
			if !hasOctopusWithHighEnergy(octopuses) {
				break
			}

			for y := 1; y < len(octopuses)-1; y++ {
				for x := 1; x < len(octopuses[y])-1; x++ {
					if octopuses[y][x] < 10 {
						continue
					}

					// Set to 0:
					octopuses[y][x] = 0

					if octopuses[y-1][x-1] > 0 {
						octopuses[y-1][x-1]++
					}
					if octopuses[y-1][x] > 0 {
						octopuses[y-1][x]++
					}
					if octopuses[y-1][x+1] > 0 {
						octopuses[y-1][x+1]++
					}
					if octopuses[y][x-1] > 0 {
						octopuses[y][x-1]++
					}
					if octopuses[y][x+1] > 0 {
						octopuses[y][x+1]++
					}
					if octopuses[y+1][x-1] > 0 {
						octopuses[y+1][x-1]++
					}
					if octopuses[y+1][x] > 0 {
						octopuses[y+1][x]++
					}
					if octopuses[y+1][x+1] > 0 {
						octopuses[y+1][x+1]++
					}
				}
			}
		}
	}
}

func allOctopusesFlashed(octopuses [12][12]int) bool {
	for y := 1; y < len(octopuses)-1; y++ {
		for x := 1; x < len(octopuses[y])-1; x++ {
			if octopuses[y][x] != 0 {
				return false
			}
		}
	}

	return true
}

func hasOctopusWithHighEnergy(octopuses [12][12]int) bool {
	for y := 1; y < len(octopuses)-1; y++ {
		for x := 1; x < len(octopuses[y])-1; x++ {
			if octopuses[y][x] > 9 {
				return true
			}
		}
	}

	return false
}

func loadOctopuses() [12][12]int {
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

	// 10x10 + border, set to -1
	var octopuses [12][12]int
	for y := 0; y < len(octopuses); y++ {
		for x := 0; x < len(octopuses[y]); x++ {
			octopuses[y][x] = -1
		}
	}

	y := 1
	for scanner.Scan() {

		for x, number := range scanner.Text() {
			energy, perr := strconv.ParseInt(string(number), 10, 8)
			if perr != nil {
				panic(perr)
			}
			octopuses[y][x+1] = int(energy)
		}

		y++
	}

	return octopuses
}
