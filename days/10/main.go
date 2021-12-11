package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var expectedOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func main() {
	lines := loadLines()
	part1(lines)
	part2(lines)
}

func part1(lines []string) int {
	errorScore := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	var score int

lineLoop:
	for _, line := range lines {

		var lastOpened []rune
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				lastOpened = append(lastOpened, char)
			case ')', ']', '}', '>':
				expected := expectedOpener[char]
				lastOpener := lastOpened[len(lastOpened)-1]
				if lastOpener != expected {
					score += errorScore[char]
					continue lineLoop
				}

				lastOpened = lastOpened[0 : len(lastOpened)-1]
			}
		}
	}

	fmt.Printf("score: %v\n", score)
	return score
}

func part2(lines []string) int {
	var scores []int

lineLoop:
	for _, line := range lines {

		var lastOpened []rune
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				lastOpened = append(lastOpened, char)
			case ')', ']', '}', '>':
				expected := expectedOpener[char]
				lastOpener := lastOpened[len(lastOpened)-1]
				if lastOpener != expected {
					continue lineLoop
				}

				lastOpened = lastOpened[0 : len(lastOpened)-1]
			}
		}

		if len(lastOpened) > 0 {
			score := calcCompletionScore(lastOpened)
			scores = append(scores, score)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	score := scores[len(scores)/2]
	fmt.Printf("%v\n", score)

	return score
}

func calcCompletionScore(opened []rune) int {
	expectedCloser := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	closerScore := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var score int

	// iterate from the last to the first, right to left:
	for i := len(opened) - 1; i >= 0; i-- {
		opener := opened[i]
		closer := expectedCloser[opener]
		score = score*5 + closerScore[closer]
	}

	return score
}

func loadLines() []string {
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

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
