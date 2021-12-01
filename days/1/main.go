package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depths := readDepths()
	part1(depths)
	part2(depths)
}

func part1(depths []int64) {
	var count int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}

	fmt.Printf("part 1: %d\n", count) // 1832
}

func part2(depths []int64) {
	const window = 3

	sum := func(values []int64) int64 {
		var sum int64

		for i := range values {
			sum += values[i]
		}

		return sum
	}

	var (
		total     = len(depths)
		to, count int
	)

	for from := 1; from < total; from++ {
		to = from + window
		if to > total {
			// Stop when there aren't enough measurements left to create a new three-measurement sum.
			break
		}

		if sum(depths[from:to]) > sum(depths[from-1:to-1]) {
			count++
		}
	}

	fmt.Printf("part 2: %d\n", count) // 1858
}

func readDepths() []int64 {
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

	var depths []int64
	for scanner.Scan() {
		depth, perr := strconv.ParseInt(scanner.Text(), 10, 64)
		if perr != nil {
			fmt.Printf("read file: %v\n", perr)
			os.Exit(1)
		}

		depths = append(depths, depth)
	}

	return depths
}
