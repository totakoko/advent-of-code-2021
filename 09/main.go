package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	fmt.Println("# Part 1", part1(input))
	fmt.Println("# Part 2", part2(input))
}

func part1(input []string) int {
	grid := parseInput(input)
	lowestPoints := findLowestPoints(grid)
	riskLevel := 0
	for _, lowestPoint := range lowestPoints {
		riskLevel += grid[lowestPoint[0]][lowestPoint[1]] + 1
	}
	return riskLevel
}

func part2(input []string) int {
	grid := parseInput(input)
	lowestPoints := findLowestPoints(grid)

	basinSizes := []int{}
	for _, lowestPoint := range lowestPoints {
		y := lowestPoint[0] // the padding added one coordinate
		x := lowestPoint[1]
		basin := []int{y*len(grid[0]) + x}
		if y > 1 {
			growBasin(grid, y-1, x, &basin, Up)
		}
		if y < len(grid)-1 {
			growBasin(grid, y+1, x, &basin, Down)
		}
		if x > 1 {
			growBasin(grid, y, x-1, &basin, Left)
		}
		if x < len(grid[0])-1 {
			growBasin(grid, y, x+1, &basin, Right)
		}
		fmt.Printf("Found basin of size %d at %d,%d\n", len(basin), y, x)
		basinSizes = append(basinSizes, len(basin))
	}

	sort.Ints(basinSizes)
	topThreeSizes := basinSizes[len(basinSizes)-3:]
	mult := 1
	for _, size := range topThreeSizes {
		mult *= size
	}
	return mult
}

func parseInput(input []string) [][]int {
	grid := make([][]int, 0, len(input)+2)
	emptyLine := make([]int, len(input[0])+2)
	for i := range emptyLine {
		emptyLine[i] = math.MaxInt64
	}
	grid = append(grid, emptyLine)
	for _, lineStr := range input {
		line := make([]int, 0, len(lineStr)+2)
		line = append(line, math.MaxInt64)
		for _, columnStr := range lineStr {
			number, _ := strconv.Atoi(string(columnStr))
			line = append(line, number)
		}
		line = append(line, math.MaxInt64)
		grid = append(grid, line)
	}
	grid = append(grid, emptyLine)
	return grid
}

func findLowestPoints(grid [][]int) [][2]int {
	lowestPoints := [][2]int{}
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			height := grid[y][x]
			if height < grid[y-1][x] && height < grid[y][x-1] && height < grid[y+1][x] && height < grid[y][x+1] {
				fmt.Printf("Found lowest point at %d,%d : %d\n", y-1, x-1, height)
				lowestPoints = append(lowestPoints, [2]int{y, x})
			}
		}
	}
	return lowestPoints
}

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

var neighborsByDirection = map[Direction][][2]int{
	Up: {
		{+1, 0}, // origin
		{-1, 0},
		{0, -1},
		{0, 1},
	},
	Down: {
		{-1, 0}, // origin
		{+1, 0},
		{0, -1},
		{0, 1},
	},
	Left: {
		{0, +1}, // origin
		{0, -1},
		{-1, 0},
		{+1, 0},
	},
	Right: {
		{0, -1}, // origin
		{0, 1},
		{-1, 0},
		{+1, 0},
	},
}

func growBasin(grid [][]int, y int, x int, basin *[]int, direction Direction) {
	if contains(*basin, y*len(grid[0])+x) {
		return
	}
	cellHeight := grid[y][x]
	neighbors := neighborsByDirection[direction]
	origin := neighbors[0]
	unvisitedNeighbors := neighbors[1:]

	originHeight := grid[y+origin[0]][x+origin[1]]
	if originHeight > cellHeight || cellHeight >= 9 {
		return
	}

	validBasin := true
	for _, neighbor := range unvisitedNeighbors {
		neighborY := y + neighbor[0]
		neighborX := x + neighbor[1]
		if neighborY > 0 && neighborX > 0 && neighborY < len(grid)+2 && neighborY < len(grid[0])+2 && cellHeight > grid[neighborY][neighborX] && !contains(*basin, neighborY*len(grid[0])+neighborX) {
			// not a basin
			validBasin = false
			break
		}
	}
	if validBasin {
		*basin = append(*basin, y*len(grid[0])+x)
		if y > 1 {
			growBasin(grid, y-1, x, basin, Up)
		}
		if y < len(grid)-1 {
			growBasin(grid, y+1, x, basin, Down)
		}
		if x > 1 {
			growBasin(grid, y, x-1, basin, Left)
		}
		if x < len(grid[0])-1 {
			growBasin(grid, y, x+1, basin, Right)
		}
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
