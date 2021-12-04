// Since none of the values in the input file are larger than 127 I decided to use int8 as type for the numbers.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers, boards := loadGameData()
	part1(numbers, boards)
}

func part1(numbers []int8, boards []board) int {
	// The basic strategy is to iterate the boards in the most stupid way possible without any regard for optimisation
	// and setting all the fields that contain the called number to 0. Setting a field to zero is our "marker", that
	// this field was "marked".
	//
	// Next, every time a field is marked the row and column are checked, if they are now have a total sum of 0. If
	// this is true, we know that either the row or column has all fields "marked". This is when we can stop iterating
	// and start creating the sum of all other fields. This is very easy at that point, because all the "marked" fields
	// have a value of zero. That's why we just need to sum up all the fields in the board.

	for _, calledNumber := range numbers {
		for _, currentBoard := range boards {
			for rowIndex, row := range currentBoard {
				for columnIndex, columnValue := range row {
					if columnValue == calledNumber {
						currentBoard[rowIndex][columnIndex] = 0

						if rowSumIsZero(row) || columnSumIsZero(currentBoard, columnIndex) {
							sum := boardSum(currentBoard)
							result := sum * int(calledNumber)
							fmt.Printf("board sum: %d\n", result)
							return result
						}
					}
				}
			}
		}
	}

	return 0
}

func part2(numbers []int8, boards []board) int {
	// In part 2 we just keep a list of the boards that already won. Once all we mark the last board as won in that
	// list, we just have to take that board, calculate the sum and multiply with the called number.
	boardWinStatus := make([]bool, len(boards))

	noLoserLeft := func(winners []bool) bool {
		for _, winner := range winners {
			if !winner {
				return false
			}
		}
		return true
	}

	for _, calledNumber := range numbers {
		for boardIndex, currentBoard := range boards {
			for rowIndex, row := range currentBoard {
				for columnIndex, columnValue := range row {
					if columnValue == calledNumber {
						currentBoard[rowIndex][columnIndex] = 0

						if rowSumIsZero(row) || columnSumIsZero(currentBoard, columnIndex) {
							boardWinStatus[boardIndex] = true

							// If there are no losers left, then the current board is the last board to win. Let's do
							// the math then :D
							if noLoserLeft(boardWinStatus) {
								sum := boardSum(currentBoard)
								result := sum * int(calledNumber)
								fmt.Printf("board sum: %d\n", result)
								return result
							}
						}
					}
				}
			}
		}
	}

	return 0
}

func boardSum(b board) int {
	var sum int
	for _, row := range b {
		for _, column := range row {
			sum += int(column)
		}
	}
	return sum
}

func rowSumIsZero(row []int8) bool {
	for i := range row {
		if row[i] > 0 {
			return false
		}
	}

	return true
}

func columnSumIsZero(b board, column int) bool {
	for rowIndex := range b {
		if b[rowIndex][column] > 0 {
			return false
		}
	}

	return true
}

// Board is a alias for a 2D array of numbers.
type board [][]int8

func loadGameData() ([]int8, []board) {
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

	scanner.Scan()                // advance scanner by one line (to the first line)
	numbersLine := scanner.Text() // read that first line
	var numbers []int8
	for _, number := range strings.Split(numbersLine, ",") {
		n, perr := strconv.ParseInt(number, 10, 8)
		if perr != nil {
			panic(perr)
		}
		numbers = append(numbers, int8(n))
	}

	// At this point the scanners position is the first line, from here on there are only boards to be read from the
	// file. Boards are "divided" by empty lines, so we need to take those into account when creating a board.
	// To make the following reading a bit simpler, advance the scanner by another line, so that it is set to the
	// first line of the first board. This is not very generic and will break if the file format is changed, but for
	// the purpose of this "exercise" it is good enough.
	scanner.Scan()

	var boards []board
	var currentBoard board
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // look for the empty lines dividing the boards
			boards = append(boards, currentBoard)
			currentBoard = nil
			continue
		}

		var row []int8
		for _, number := range strings.Split(line, " ") {
			// Due to the formatting of the file single digit numbers are indented by another space. So splitting at
			// "space" will create empty strings that should be skipped.
			if number == "" {
				continue
			}
			p, perr := strconv.ParseInt(number, 10, 8)
			if perr != nil {
				panic(perr)
			}
			row = append(row, int8(p))
		}
		currentBoard = append(currentBoard, row)
	}

	return numbers, boards
}
