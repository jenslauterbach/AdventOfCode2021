package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	report := loadReport()
	part1(report)
	part2(report)
}

func part1(report []uint64) uint64 {
	var (
		// n is the number of bits that are considered reading from right to left. Since we are using uint64 there are a
		// lot of unused bits that we should consider.
		n = 12
		// op is used to determine if the bit at position 0 is a one or zero by using bitwise AND (&) with the current
		// diagnostic from the report.
		op uint64 = 1
		// gamma, epsilon and (power) consumption are the values we are looking for.
		gamma, epsilon, consumption uint64
		// ones is a counter for the number of times the bit at position 0 is a one or zero.
		ones uint64
		// threshold is the number above which the count of ones has to be, to be in the majority. So if the report
		// contains 1_000 diagnostics, the number one has to appear at least 501 one times to be in the majority.
		threshold = uint64(len(report)/2) + 1
	)

	// Iterate over every bit that is considered one at a time, starting with the least significant bit (position 0) or
	// the "right most bit".
	//
	// position:  321076543210
	// bits:      000000000000
	// start:                ^
	for i := 0; i < n; i++ {
		// For every bit position that we want to consider, we have to iterate every diagnostic report and check the
		// bit at the currently observed position (i).
		for _, diagnostic := range report {
			// Every report is shifted "i" times. For the first iteration that means "diagnostic" is shifted 0 times to
			// the right and when "i" = 11, diagnostic is shifted eleven times. This is done to move to currently
			// observed bit to the right most position. After that we bitwise AND the remaining value with 1 to set all
			// bits to 0 except for the bit at position 1.
			//
			// Example:
			//
			// observed bit position (i): 2
			// diagnostic:                1000
			// diagnostic >> 2:           0010 (moved the 1 to the right 2 times)
			// (diagnostic >> 2) & 1      0000 (bitwise AND of 0010 & 0001 = 0000)
			// ---------------------------------------------------------------------------------------------------------
			// observed bit position (i): 3
			// diagnostic:                1000
			// diagnostic >> 3:           0001 (moved the 1 to the right 3 times)
			// (diagnostic >> 3) & 1      0001 (bitwise AND of 0001 & 0001 = 0001)
			//
			// This means that the remaining value is either a "0" or "1". This remaining value is then just added to
			// the counter called "ones", either increasing its value (1) or not (0).
			ones += (diagnostic >> i) & op
		}

		// Only if the number of "ones" is higher than the threshold we want gamma to be "updated". But we need to make
		// sure that it is updated at the right "bit" position. The position is determined by "i" or our currently
		// observed bit. So if we currently observe bit 4 (i=3) we bitwise OR the current gamma value with the
		// appropriate bit mask that would set the bit at that position to 1.
		if ones > threshold {
			gamma |= 1 << i
		}

		ones = 0 // reset the number of ones
	}

	// epsilon is just the bitwise "inverse" of the gamma value. So we just "flip" the bits with the XOR operand. Since
	// we just want to observe 12 bits overall, we XOR with the binary 0000111111111111 (or 0x0ffff).
	epsilon = gamma ^ 0x0fff
	consumption = gamma * epsilon

	fmt.Printf("gamma: %012b (%d)\n", gamma, gamma)
	fmt.Printf("epsilon: %012b (%d)\n", epsilon, epsilon)
	fmt.Printf("power consumption: %d\n", gamma*epsilon)

	return consumption
}

func part2(report []uint64) uint64 {
	var op uint64 = 1

	findOxygenRating := func(reports []uint64) uint64 {
		remainingReports := report
		for i := 11; i >= 0 && len(remainingReports) >= 1; i-- {
			var zeros, ones []uint64

			for _, diagnostic := range remainingReports {
				bit := (diagnostic >> i) & op

				switch bit {
				case 0:
					zeros = append(zeros, diagnostic)
				case 1:
					ones = append(ones, diagnostic)
				}
			}

			if len(zeros) > len(ones) {
				remainingReports = zeros
			} else {
				remainingReports = ones
			}
		}

		return remainingReports[0]
	}

	findCO2Rating := func(reports []uint64) uint64 {
		remainingReports := report
		for i := 11; i >= 0 && len(remainingReports) >= 1; i-- {
			var zeros, ones []uint64

			for _, diagnostic := range remainingReports {
				bit := (diagnostic >> i) & op

				switch bit {
				case 0:
					zeros = append(zeros, diagnostic)
				case 1:
					ones = append(ones, diagnostic)
				}
			}

			switch {
			case len(zeros) == 0:
				remainingReports = ones
			case len(ones) == 0:
				remainingReports = zeros
			case len(zeros) > len(ones):
				remainingReports = ones
			case len(zeros) <= len(ones):
				remainingReports = zeros
			}
		}

		return remainingReports[0]
	}

	oxygenRating := findOxygenRating(report)
	CO2Rating := findCO2Rating(report)
	lifeSupportRating := oxygenRating * CO2Rating

	fmt.Printf("oxygen rating: %12b (%d)\n", oxygenRating, oxygenRating)
	fmt.Printf("CO2 rating: %12b (%d)\n", CO2Rating, CO2Rating)
	fmt.Printf("life support rating: %d\n", lifeSupportRating)

	return 0
}

func loadReport() []uint64 {
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

	var report []uint64
	for scanner.Scan() {
		p, perr := strconv.ParseUint(scanner.Text(), 2, 12)
		if perr != nil {
			panic(perr)
		}
		report = append(report, p)
	}

	//fmt.Printf("%b", 1<<11)

	return report
}
