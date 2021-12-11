package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

const gridSize = 10

type Grid [][]int

func part1(input []string) int {
	grid := parseInput(input)
	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += step(grid)
	}
	return totalFlashes
}

// parse the input and wrap the grid with a padding line/column of MinInt64
// so that these cells will never flash
func parseInput(input []string) Grid {
	grid := make([][]int, 0, len(input)+2)
	emptyLine := make([]int, len(input[0])+2)
	for i := range emptyLine {
		emptyLine[i] = math.MinInt64
	}
	grid = append(grid, emptyLine)
	for _, lineStr := range input {
		line := make([]int, 0, len(lineStr)+2)
		line = append(line, math.MinInt64)
		for _, columnStr := range lineStr {
			number, _ := strconv.Atoi(string(columnStr))
			line = append(line, number)
		}
		line = append(line, math.MinInt64)
		grid = append(grid, line)
	}
	grid = append(grid, emptyLine)
	return grid
}

func step(grid Grid) int {
	for y := 1; y < gridSize+1; y++ {
		for x := 1; x < gridSize+1; x++ {
			energy := &grid[y][x]
			*energy++
		}
	}
	return stepFlashes(grid)
}

func stepFlashes(grid Grid) int {
	flashesCount := 0
	hasFlashed := true
	for hasFlashed {
		hasFlashed = false
		for y := 1; y < gridSize+1; y++ {
			for x := 1; x < gridSize+1; x++ {
				energy := &grid[y][x]
				if *energy > 9 {
					// flash the current octopus
					flashesCount++
					*energy = 0

					// increase neighbors energy only if > 0
					stepOctopus(&grid[y-1][x-1])
					stepOctopus(&grid[y-1][x])
					stepOctopus(&grid[y-1][x+1])
					stepOctopus(&grid[y][x-1])
					stepOctopus(&grid[y][x+1])
					stepOctopus(&grid[y+1][x-1])
					stepOctopus(&grid[y+1][x])
					stepOctopus(&grid[y+1][x+1])

					// will trigger a new detection loop
					hasFlashed = true
				}
			}
		}
	}
	return flashesCount
}

func stepOctopus(energy *int) {
	if *energy > 0 {
		*energy++
	}
}

func part2(input []string) int {
	grid := parseInput(input)
	for i := 0; i < 1000000; i++ {
		totalFlashes := step(grid)
		if totalFlashes == gridSize*gridSize {
			return i + 1
		}
	}
	return 0
}

func printGrid(grid Grid) {
	for y := 1; y < gridSize+1; y++ {
		for x := 1; x < gridSize+1; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}
