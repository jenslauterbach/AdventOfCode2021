package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := loadLines()
	part1(lines)
	part2(lines)
}

func part1(lines []line) int {
	maxX, maxY := findMaxima(lines)
	board := make([][]int, maxX+1)
	for x := range board {
		board[x] = make([]int, maxY+1)
	}

	// Iterate all the lines and draw the horizontal and vertical lines onto the board. Keep in mind, that lines can go
	// in two directions (left <-> right or up <-> down). Therefore, when iterating those lines when drawing them onto
	// the board might require a change of "iteration direction".
	for _, currentLine := range lines {
		switch {
		case currentLine.isHorizontal():
			y, startX, endX := currentLine.start.y, currentLine.start.x, currentLine.end.x

			// consider lines going from right to left
			if currentLine.start.x > currentLine.end.x {
				startX, endX = currentLine.end.x, currentLine.start.x
			}

			for x := startX; x <= endX; x++ {
				board[x][y] = board[x][y] + 1
			}
		case currentLine.isVertical():
			x, startY, endY := currentLine.start.x, currentLine.start.y, currentLine.end.y

			// consider lines going from bottom to top
			if currentLine.start.y > currentLine.end.y {
				startY, endY = currentLine.end.y, currentLine.start.y
			}

			for y := startY; y <= endY; y++ {
				board[x][y] = board[x][y] + 1
			}
		}
	}

	overlappingPoints := numberOfOverlappingPoints(board)
	fmt.Printf("overlapping points: %d\n", overlappingPoints)

	return overlappingPoints
}

func part2(lines []line) int {
	maxX, maxY := findMaxima(lines)
	board := make([][]int, maxX+1)
	for x := range board {
		board[x] = make([]int, maxY+1)
	}

	// Iterate all the lines and draw the horizontal and vertical lines onto the board. Keep in mind, that lines can go
	// in two directions (left <-> right or up <-> down). Therefore, when iterating those lines when drawing them onto
	// the board might require a change of "iteration direction".
	for _, currentLine := range lines {
		switch {
		case currentLine.isHorizontal():
			y, startX, endX := currentLine.start.y, currentLine.start.x, currentLine.end.x

			// consider lines going from right to left
			if currentLine.start.x > currentLine.end.x {
				startX, endX = currentLine.end.x, currentLine.start.x
			}

			for x := startX; x <= endX; x++ {
				board[x][y] = board[x][y] + 1
			}
		case currentLine.isVertical():
			x, startY, endY := currentLine.start.x, currentLine.start.y, currentLine.end.y

			// consider lines going from bottom to top
			if currentLine.start.y > currentLine.end.y {
				startY, endY = currentLine.end.y, currentLine.start.y
			}

			for y := startY; y <= endY; y++ {
				board[x][y] = board[x][y] + 1
			}
		default: // diagonal line
			// lr = left to right
			// rl = right to left
			// tb = top to bottom
			// bt = bottom to top
			// only one of the following can be true per line:
			lrtb := (currentLine.start.x < currentLine.end.x) && (currentLine.start.y < currentLine.end.y)
			lrbt := (currentLine.start.x < currentLine.end.x) && (currentLine.start.y > currentLine.end.y)
			rlbt := (currentLine.start.x > currentLine.end.x) && (currentLine.start.y > currentLine.end.y)
			rltb := (currentLine.start.x > currentLine.end.x) && (currentLine.start.y < currentLine.end.y)

			startX := currentLine.start.x
			startY := currentLine.start.y
			endX := currentLine.end.x
			endY := currentLine.end.y

			switch {
			case lrtb:
				for x, y := startX, startY; x <= endX && y <= endY; x, y = x+1, y+1 {
					board[x][y] = board[x][y] + 1
				}
			case lrbt:
				for x, y := startX, startY; x <= endX && y >= endY; x, y = x+1, y-1 {
					board[x][y] = board[x][y] + 1
				}
			case rlbt:
				for x, y := startX, startY; x >= endX && y >= endY; x, y = x-1, y-1 {
					board[x][y] = board[x][y] + 1
				}
			case rltb:
				for x, y := startX, startY; x >= endX && y <= endY; x, y = x-1, y+1 {
					board[x][y] = board[x][y] + 1
				}
			}
		}
	}

	overlappingPoints := numberOfOverlappingPoints(board)
	fmt.Printf("overlapping points: %d\n", overlappingPoints)

	return overlappingPoints
}

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func (l line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l line) isVertical() bool {
	return l.start.x == l.end.x
}

func loadLines() []line {
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

	parsePoint := func(rawPoint string) point {
		coords := strings.Split(rawPoint, ",")
		startX, perr := strconv.ParseInt(coords[0], 10, 16)
		if perr != nil {
			panic(perr)
		}

		startY, perr := strconv.ParseInt(coords[1], 10, 16)
		if perr != nil {
			panic(perr)
		}

		return point{
			x: int(startX),
			y: int(startY),
		}
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var lines []line
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")

		start := parsePoint(coords[0])
		end := parsePoint(coords[1])

		newLine := line{
			start: start,
			end:   end,
		}

		lines = append(lines, newLine)
	}

	return lines
}

func findMaxima(lines []line) (maxX, maxY int) {
	for _, l := range lines {
		if l.start.x > maxX {
			maxX = l.start.x
		}
		if l.end.x > maxX {
			maxX = l.end.x
		}
		if l.start.y > maxY {
			maxY = l.start.y
		}
		if l.end.y > maxY {
			maxY = l.end.y
		}
	}
	return
}

func numberOfOverlappingPoints(board [][]int) int {
	var overlappingPoints int
	for x := range board {
		for y := range board[x] {
			if board[x][y] > 1 {
				overlappingPoints++
			}
		}
	}

	return overlappingPoints
}
