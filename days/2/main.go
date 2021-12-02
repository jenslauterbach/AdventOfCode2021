package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands := readCommands()
	part1(commands)
	part2(commands)
}

func part1(commands []command) {
	var x, depth int
	for _, cmd := range commands {
		switch cmd.direction {
		case forward:
			x += cmd.amount
		case down:
			depth += cmd.amount
		case up:
			depth -= cmd.amount
			if depth < 0 {
				depth = 0
			}
		}
	}

	fmt.Printf("x: %d, depth: %d, product: %d\n", x, depth, x*depth)
}

func part2(commands []command) {
	var x, depth, aim int
	for _, cmd := range commands {
		switch cmd.direction {
		case forward:
			x += cmd.amount
			depth += aim * cmd.amount
			if depth < 0 {
				depth = 0
			}
		case down:
			aim += cmd.amount
		case up:
			aim -= cmd.amount
		}
	}

	fmt.Printf("x: %d, depth: %d, product: %d\n", x, depth, x*depth)
}

type direction string

const (
	forward direction = "forward"
	up      direction = "up"
	down    direction = "down"
)

type command struct {
	direction direction
	amount    int
}

func readCommands() []command {
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

	parseAction := func(line string) direction {
		c := strings.Split(line, " ")[0]
		switch c {
		case "forward":
			return forward
		case "down":
			return down
		case "up":
			return up
		default:
			panic(fmt.Sprintf("unknown command: %v", c))
		}
	}

	parseAmount := func(line string) int {
		c := strings.Split(line, " ")[1]
		p, perr := strconv.ParseInt(c, 10, 8)
		if perr != nil {
			panic(perr)
		}

		return int(p)
	}

	var commands []command
	for scanner.Scan() {
		line := scanner.Text()
		cmd := command{
			direction: parseAction(line),
			amount:    parseAmount(line),
		}

		commands = append(commands, cmd)
	}

	return commands
}
