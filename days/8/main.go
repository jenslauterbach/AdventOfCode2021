package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	displays := loadDisplays()
	part1(displays)
	part2(displays)
}

func part1(displays []display) {
	var count int
	for _, d := range displays {
		for _, o := range d.outputs {
			n := len(o)
			switch n {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	fmt.Printf("number of 1,4,7,8: %d\n", count)
}

func part2(displays []display) {
	digits := map[int]int{
		36: 0,
		8:  1,
		10: 2,
		30: 3,
		32: 4,
		15: 5,
		18: 6,
		12: 7,
		56: 8,
		48: 9,
	}

	var count int

	for _, d := range displays {
		one := patternsWithLength(d.patterns, 2)
		four := patternsWithLength(d.patterns, 4)
		eight := patternsWithLength(d.patterns, 7)

		var number int

		for i, output := range d.outputs {
			product := productOfSharedSegments([]rune(output), one[0], four[0], eight[0])
			digit := digits[product]

			switch i {
			case 0:
				number += digit * 1000
			case 1:
				number += digit * 100
			case 2:
				number += digit * 10
			case 3:
				number += digit * 1
			}
		}

		count += number
	}

	fmt.Printf("count: %d\n", count)
}

type display struct {
	patterns []string
	outputs  []string
}

func loadDisplays() []display {
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

	var displays []display

	for scanner.Scan() {
		line := scanner.Text()

		displays = append(displays, display{
			patterns: parsePatterns(line),
			outputs:  parseOutput(line),
		})
	}

	return displays
}

func parsePatterns(line string) []string {
	rawPatterns := strings.Split(line, " | ")[0]

	var patterns []string
	for _, pattern := range strings.Split(rawPatterns, " ") {
		patterns = append(patterns, pattern)
	}

	return patterns
}

func parseOutput(line string) []string {
	rawOutput := strings.Split(line, " | ")[1]

	var digits []string
	for _, digit := range strings.Split(rawOutput, " ") {
		digits = append(digits, digit)
	}

	return digits
}

func difference(a, b []rune) []rune {
	var diff []rune

	for _, aRune := range a {
		var contained bool

		for _, bRune := range b {
			if aRune == bRune {
				contained = true
				break
			}
		}

		if !contained {
			diff = append(diff, aRune)
		}
	}

	return diff
}

func patternsWithLength(patterns []string, length int) [][]rune {
	var foundPatterns []string
	for _, pattern := range patterns {
		if len(pattern) == length {
			foundPatterns = append(foundPatterns, pattern)
		}
	}

	var result [][]rune
	for _, p := range foundPatterns {
		var res []rune
		for _, r := range p {
			res = append(res, r)
		}
		result = append(result, res)
	}

	return result
}

func subtract(a []rune, b ...[]rune) []rune {
	var result []rune

	for _, aRune := range a {
		for _, bRunes := range b {
			var contained bool
			for _, bRune := range bRunes {
				if aRune == bRune {
					contained = true
					break
				}
			}
			if !contained {
				result = append(result, aRune)
			}
		}
	}

	return result
}

func productOfSharedSegments(a, one, four, eight []rune) int {
	n := len(a)
	sharedOne := n - len(difference(a, one))
	sharedFour := n - len(difference(a, four))
	sharedEight := n - len(difference(a, eight))

	return sharedOne * sharedFour * sharedEight
}
