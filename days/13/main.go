package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dots, folds := loadPage()

	count := part1(dots, folds)
	fmt.Printf("part1: %d\n", count)

	part2(dots, folds)
}

func part1(dots []dot, folds []fold) int {
	f := folds[0]

	var folder foldFunc

	switch f.axis {
	case X:
		folder = foldLeft
	case Y:
		folder = foldUp
	}

	dots = foldDots(dots, folder, f.amount)

	return numberOfUniqueDots(dots)
}

func part2(dots []dot, folds []fold) {
	for _, f := range folds {
		var folder foldFunc

		switch f.axis {
		case X:
			folder = foldLeft
		case Y:
			folder = foldUp
		}

		dots = foldDots(dots, folder, f.amount)
	}

	displayDots(dots)
}

func loadPage() ([]dot, []fold) {
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

	var dots []dot
	var folds []fold

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.TrimSpace(line) == "":
			break
		case strings.Contains(line, "fold"):
			coords := strings.TrimPrefix(line, "fold along ")
			split := strings.Split(coords, "=")

			var foldAxis axis
			switch split[0] {
			case "x":
				foldAxis = X
			case "y":
				foldAxis = Y
			}

			amount, perr := strconv.ParseInt(split[1], 10, 16)
			if perr != nil {
				panic(perr)
			}

			folds = append(folds, fold{
				axis:   foldAxis,
				amount: int(amount),
			})
		default:
			coords := strings.Split(line, ",")

			x, perr := strconv.ParseInt(coords[0], 10, 16)
			if perr != nil {
				panic(perr)
			}

			y, perr := strconv.ParseInt(coords[1], 10, 16)
			if perr != nil {
				panic(perr)
			}

			dots = append(dots, dot{
				x: int(x),
				y: int(y),
			})
		}
	}

	return dots, folds
}

func displayDots(dots []dot) {
	var maxX, maxY int
	for _, d := range dots {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}

	// initialize display matrix
	display := make([][]bool, maxY+1)
	for i := range display {
		display[i] = make([]bool, maxX+1)
	}

	// set dots in display matrix to true
	for _, d := range dots {
		display[d.y][d.x] = true
	}

	// iterate display matrix and display only the dots that are true as #, the rest as simple dot. Monospace fonts
	// should make sure that everything is aligned properly.
	for _, row := range display {
		for _, column := range row {
			if column {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func numberOfUniqueDots(dots []dot) int {
	// This function abuses the fact that 'dot' struct can be used as a map key. That means that if two dots with the
	// same x,y are added (they are the same), they will only appear once in the map. Since keys are "unique". After
	// all dots are added to the map, all that needs to be done is to check the length of the map to find out how many
	// of the dots have been unique.
	page := make(map[dot]bool)

	for _, d := range dots {
		page[d] = true
	}

	return len(page)
}

type foldFunc func(x, y, foldCoordinate int) (int, int)

func foldUp(x, y, foldCoordinate int) (int, int) {
	// do nothing if the current y is smaller (above) than the foldCoordinate.
	if y < foldCoordinate {
		return x, y
	}

	distance := y - foldCoordinate

	return x, foldCoordinate - distance
}

func foldLeft(x, y, foldCoordinate int) (int, int) {
	// do nothing if the current x is smaller (to the left) than the foldCoordinate.
	if x < foldCoordinate {
		return x, y
	}

	distance := x - foldCoordinate

	return foldCoordinate - distance, y
}

func foldDots(dots []dot, foldFunc foldFunc, foldCoordinate int) []dot {
	var foldedDots []dot

	for _, d := range dots {
		x, y := foldFunc(d.x, d.y, foldCoordinate)
		foldedDots = append(foldedDots, dot{
			x: x,
			y: y,
		})
	}

	return foldedDots
}

type axis int

const (
	X axis = iota
	Y
)

type fold struct {
	axis   axis
	amount int
}

type dot struct {
	x int
	y int
}

func (d dot) String() string {
	return fmt.Sprintf("[x: %d, y: %d]", d.x, d.y)
}
