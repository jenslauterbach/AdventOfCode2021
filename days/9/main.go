package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	m := loadHeightMap()
	part1(m)
	part2(m)
}

func part1(m [][]int) int {
	points := findLowPoints(m)

	var risk int
	for _, p := range points {
		risk += 1 + p.height
	}

	fmt.Printf("risk: %d\n", risk)

	return risk
}

func part2(m [][]int) int {
	lowPoints := findLowPoints(m)

	var basinSizes []int
	for _, lowPoint := range lowPoints {
		b := basin{included: map[point]bool{lowPoint: true}}
		buildBasin(m, &b)
		basinSizes = append(basinSizes, b.size())
	}

	// find three largest basins by sorting descending and then getting the sub slice of the top 3 elements.
	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	largest := basinSizes[0:3]
	result := largest[0] * largest[1] * largest[2]

	fmt.Printf("%d * %d * %d = %d\n", largest[0], largest[1], largest[2], result)
	return result
}

type point struct {
	x, y, height int
}

func findLowPoints(m [][]int) []point {
	var points []point

	yMax := len(m) - 1
	for y := range m {
		row := m[y]
		xMax := len(row) - 1

		for x := range row {
			height := m[y][x]
			left, right, up, down := 9, 9, 9, 9

			switch {
			// Upper, left corner:
			case y == 0 && x == 0:
				right = m[y][x+1]
				down = m[y+1][x]
			// Upper, right corner:
			case y == 0 && x == xMax:
				left = m[y][x-1]
				down = m[y+1][x]
			// Upper Edge (except corners)
			case y == 0 && x > 0 && x < xMax:
				left = m[y][x-1]
				right = m[y][x+1]
				down = m[y+1][x]
			// Left Edge (except corners):
			case y > 0 && y < yMax && x == 0:
				right = m[y][x+1]
				up = m[y-1][x]
				down = m[y+1][x]
			// Right Edge (except corners):
			case y > 0 && y < yMax && x == xMax:
				left = m[y][x-1]
				up = m[y-1][x]
				down = m[y+1][x]
			// Bottom edge:
			case y == yMax && x > 0 && x < xMax:
				left = m[y][x-1]
				right = m[y][x+1]
				up = m[y-1][x]
			// Bottom, left corner:
			case y == yMax && x == 0:
				right = m[y][x+1]
				up = m[y-1][x]
			// Bottom, right corner:
			case y == yMax && x == xMax:
				left = m[y][x-1]
				up = m[y-1][x]
			// All other locations:
			default:
				left = m[y][x-1]
				right = m[y][x+1]
				up = m[y-1][x]
				down = m[y+1][x]
			}

			if isLowPoint(height, left, right, up, down) {
				points = append(points, point{
					x:      x,
					y:      y,
					height: height,
				})
			}
		}
	}

	return points
}

func loadHeightMap() [][]int {
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

	var heightMap [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var digits []int
		for _, character := range line {
			d, perr := strconv.ParseInt(string(character), 10, 8)
			if perr != nil {
				panic(perr)
			}

			digits = append(digits, int(d))
		}

		heightMap = append(heightMap, digits)
	}

	return heightMap
}

func isLowPoint(v int, others ...int) bool {
	for _, o := range others {
		if v >= o {
			return false
		}
	}

	return true
}

func noneNineNeighbours(m [][]int, p point) []point {
	yMax := len(m) - 1
	xMax := len(m[p.y]) - 1

	var points []point

	// left
	if p.x > 0 {
		height := m[p.y][p.x-1]
		if height != 9 {
			points = append(points, point{
				x:      p.x - 1,
				y:      p.y,
				height: height,
			})
		}
	}

	// right
	if p.x < xMax {
		height := m[p.y][p.x+1]
		if height != 9 {
			points = append(points, point{
				x:      p.x + 1,
				y:      p.y,
				height: height,
			})
		}
	}

	// up
	if p.y > 0 {
		height := m[p.y-1][p.x]
		if height != 9 {
			points = append(points, point{
				x:      p.x,
				y:      p.y - 1,
				height: height,
			})
		}
	}

	// down
	if p.y < yMax {
		height := m[p.y+1][p.x]
		if height != 9 {
			points = append(points, point{
				x:      p.x,
				y:      p.y + 1,
				height: height,
			})
		}
	}

	return points
}

type basin struct {
	included map[point]bool
}

func (b *basin) size() int {
	return len(b.included)
}

func (b *basin) isIncluded(p point) bool {
	for basinPoint := range b.included {
		if basinPoint == p {
			return true
		}
	}

	return false
}

func (b *basin) include(p point) {
	b.included[p] = true
}

func buildBasin(m [][]int, b *basin) {
	for p := range b.included {
		neighbours := noneNineNeighbours(m, p)

		if len(neighbours) == 0 {
			continue
		}

		for _, neighbour := range neighbours {
			if b.isIncluded(neighbour) {
				continue
			}

			b.include(neighbour)
			buildBasin(m, b)
		}
	}
}
