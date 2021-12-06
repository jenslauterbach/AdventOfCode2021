package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fishes := loadFish()
	//part1(fishes)
	part2(fishes)
}

func part1(fishes []int8) int {
	for day := 0; day < 80; day++ {
		for i := range fishes {
			switch fishes[i] {
			case 0:
				fishes[i] = 6
				fishes = append(fishes, 8)
			default:
				fishes[i] -= 1
			}
		}
	}

	fmt.Printf("Number of fishes after 80 days: %d\n", len(fishes))

	return len(fishes)
}

// The point of this puzzle is that you can't use the simple approach from part1 because you will run out of memory.
// The slice used in part1 will grow immensely until it does not fit into memory anymore and your program is getting
// sig killed.
//
// That means we have to be more clever for part 2 and find a better strategy to calculate how many fish are born every
// day and how this is tracked.
func part2(fishes []int8) uint64 {
	// remainingDays is a slice of 9 indices. One index per "number of remaining days". So remainingDays[4] contains the
	// amount of fish that have 4 days left until the "hatch" or lay an egg ;)
	remainingDays := []uint64{
		0, // 0 days left
		0, // 1 days left
		0, // 2 days left
		0, // 3 days left
		0, // 4 days left
		0, // 5 days left
		0, // 6 days left
		0, // 7 days left
		0, // 8 days left
	}

	// Initialise the remainingDays with the fish that have been read from the input file
	for i := range fishes {
		r := fishes[i]
		remainingDays[r] += 1
	}

	// Now iterate 256 times (once per day) and do the following:
	//
	// 1. Remember how many fish have 0 days left until they "hatch" (which means they will bear a new fish today)
	// 2. Copy over all the days to their respective next days (copy the value of day 1 to 0, 2 to 1, 3 to 2 etc)
	// 3. For the 6 day, not only copy the value of the 7th day, but also add the number of "hatchling" parents
	// 4. Set day 8 to the number of newly born hatchlings.
	for day := 0; day < 256; day++ {
		hatchlings := remainingDays[0]

		remainingDays[0] = remainingDays[1]
		remainingDays[1] = remainingDays[2]
		remainingDays[2] = remainingDays[3]
		remainingDays[3] = remainingDays[4]
		remainingDays[4] = remainingDays[5]
		remainingDays[5] = remainingDays[6]
		remainingDays[6] = remainingDays[7] + hatchlings
		remainingDays[7] = remainingDays[8]
		remainingDays[8] = hatchlings
	}

	var sum uint64
	for i := range remainingDays {
		sum += remainingDays[i]
	}

	return sum
}

func loadFish() []int8 {
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

	var fishes []int8
	for _, number := range numbers {
		n, perr := strconv.ParseInt(number, 10, 8)
		if perr != nil {
			panic(perr)
		}
		fishes = append(fishes, int8(n))
	}

	return fishes
}
