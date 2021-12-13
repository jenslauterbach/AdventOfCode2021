package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	caves := loadCaves()

	numberPathsPart1 := part1(caves)
	fmt.Printf("part 1: %d\n", numberPathsPart1)

	numberPathsPart2 := part2(caves)
	fmt.Printf("part 2: %d\n", numberPathsPart2)
}

func part1(caves []*cave) int {
	start := findCave(caves, "start")
	paths := findPaths(path{caves: []*cave{start}}, false)
	return len(paths)
}

func part2(caves []*cave) int {
	start := findCave(caves, "start")
	paths := findPaths(path{caves: []*cave{start}}, true)
	return len(paths)
}

func findPaths(p path, allowDoubleVisit bool) []path {
	var paths []path

	currentCave := p.caves[len(p.caves)-1]
	if currentCave.isEnd() {
		return []path{p}
	}

	for _, neighbour := range currentCave.neighbours {
		if allowDoubleVisit && p.hasCaveThatWasVisitedTwice() && p.contains(*neighbour) && neighbour.isSmall() {
			continue
		}

		if !allowDoubleVisit && p.contains(*neighbour) && neighbour.isSmall() {
			continue
		}

		newPath := path{caves: p.caves}
		newPath.add(neighbour)

		paths = append(paths, findPaths(newPath, allowDoubleVisit)...)
	}

	return paths
}

func loadCaves() []*cave {
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

	var caves []*cave

	// Two caves per lines
	for scanner.Scan() {
		line := scanner.Text()

		caveNames := strings.Split(line, "-")

		var c1, c2 *cave

		c1 = findCave(caves, caveNames[0])
		if c1 == nil {
			c1 = &cave{name: caveNames[0]}
			caves = append(caves, c1)
		}

		c2 = findCave(caves, caveNames[1])
		if c2 == nil {
			c2 = &cave{name: caveNames[1]}
			caves = append(caves, c2)
		}

		// No cave should point to the start, because nobody should be allowed to go "back" to the start.
		// Also not set any neighbours for the end cave. Once you are at the end, you shouldn't go anywhere else.
		if !c1.isStart() && !c2.isEnd() {
			c2.addNeighbour(c1)
		}
		if !c2.isStart() && !c1.isEnd() {
			c1.addNeighbour(c2)
		}
	}

	return caves
}

type path struct {
	caves []*cave
}

func (p path) String() string {
	b := strings.Builder{}
	for i := range p.caves {
		b.WriteString(p.caves[i].name)
		b.WriteString(",")
	}

	return b.String()
}

func (p path) contains(c cave) bool {
	for i := range p.caves {
		if p.caves[i].name == c.name {
			return true
		}
	}

	return false
}

func (p *path) add(c *cave) {
	p.caves = append(p.caves, c)
}

func (p path) hasCaveThatWasVisitedTwice() bool {
	visits := make(map[string]int)
	for _, c := range p.caves {
		if c.isLarge() {
			continue
		}

		if visits[c.name] == 1 {
			return true
		}

		visits[c.name]++
	}

	return false
}

type cave struct {
	name       string
	neighbours []*cave
}

func (c cave) String() string {
	return c.name
}

func (c cave) isSmall() bool {
	return strings.ToLower(c.name) == c.name
}

func (c cave) isLarge() bool {
	return !c.isSmall()
}

func (c cave) isStart() bool {
	return c.name == "start"
}

func (c cave) isEnd() bool {
	return c.name == "end"
}

func (c *cave) addNeighbour(neighbour *cave) {
	if findCave(c.neighbours, neighbour.name) != nil {
		return
	}

	c.neighbours = append(c.neighbours, neighbour)
}

func findCave(caves []*cave, name string) *cave {
	for i := range caves {
		if caves[i].name == name {
			return caves[i]
		}
	}

	return nil
}
