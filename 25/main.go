package main

import (
	"fmt"
	"io/ioutil"
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
	return evolve(grid)
}

func part2(input []string) int {
	return 0
}

type Grid [][]Cell

func (grid Grid) Clone() Grid {
	clone := make([][]Cell, len(grid))
	for y := 0; y < len(grid); y++ {
		clone[y] = make([]Cell, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			clone[y][x] = grid[y][x]
		}
	}
	return clone
}

func (grid Grid) Print() {
	// fmt.Print("\033[H\033[2J") // clear the terminal
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			char := '.'
			switch grid[y][x] {
			case EastCucumber:
				char = '>'
			case SouthCucumber:
				char = 'v'
			}
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func NewEmptyGrid(height int, width int) Grid {
	grid := make([][]Cell, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]Cell, width)
		for x := 0; x < width; x++ {
			grid[y][x] = Empty
		}
	}
	return grid
}

type Cell uint8

const (
	Empty         Cell = 1
	EastCucumber  Cell = 2
	SouthCucumber Cell = 3
)

func parseInput(input []string) Grid {
	grid := make([][]Cell, 0, len(input))
	for _, lineStr := range input {
		line := make([]Cell, 0, len(lineStr))
		for _, cell := range lineStr {
			switch cell {
			case '>':
				line = append(line, EastCucumber)
			case 'v':
				line = append(line, SouthCucumber)
			case '.':
				line = append(line, Empty)
			default:
				panic("unknown char: " + string(cell))
			}
		}
		grid = append(grid, line)
	}
	return grid
}

func evolve(initialGrid Grid) int {
	iterations := 0
	lastVerticalIndex := len(initialGrid) - 1
	lastHorizontalIndex := len(initialGrid[0]) - 1

	var currentGrid Grid
	futureGrid := initialGrid
	atLeastOneMovement := true
	for atLeastOneMovement {
		iterations++
		atLeastOneMovement = false

		// east cucumbers
		currentGrid = futureGrid.Clone()
		for y := 0; y <= lastVerticalIndex; y++ {
			for x := 0; x <= lastHorizontalIndex; x++ {
				if currentGrid[y][x] == EastCucumber {
					nextCellIndex := x + 1
					if x == lastHorizontalIndex {
						nextCellIndex = 0
					}
					if currentGrid[y][nextCellIndex] == Empty {
						futureGrid[y][x] = Empty
						futureGrid[y][nextCellIndex] = EastCucumber
						x++
						atLeastOneMovement = true
					}
				}
			}
		}

		// south cucumbers
		currentGrid = futureGrid.Clone()
		for x := 0; x <= lastHorizontalIndex; x++ {
			for y := 0; y <= lastVerticalIndex; y++ {
				if currentGrid[y][x] == SouthCucumber {
					nextCellIndex := y + 1
					if y == lastVerticalIndex {
						nextCellIndex = 0
					}
					if currentGrid[nextCellIndex][x] == Empty {
						futureGrid[y][x] = Empty
						futureGrid[nextCellIndex][x] = SouthCucumber
						y++
						atLeastOneMovement = true
					}
				}
			}
		}
	}
	return iterations
}
