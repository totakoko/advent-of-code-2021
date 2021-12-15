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

type Cell struct {
	Id        int
	X         int
	Y         int
	TotalRisk int
}

func part1(input []string) int {
	grid := parseInput(input)
	return findShortestPathRisk(grid)
}

func part2(input []string) int {
	grid := parseBiggerInput(input)
	return findShortestPathRisk(grid)
}

var cellNeighbors = [][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{+1, 0},
}

// Parse the grid and add a padding line/column
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

func findShortestPathRisk(grid [][]int) int {
	gridWidth := len(grid[0])

	targetX := gridWidth - 2
	targetY := len(grid) - 2

	cellsToVisit := []Cell{
		{
			Id:        gridWidth,
			X:         1,
			Y:         1,
			TotalRisk: 0, // first cell does not count
		},
	}
	cellsTotalRisks := make(map[int]int, gridWidth*len(grid))
	for i := 0; i < gridWidth*len(grid); i++ {
		cellsTotalRisks[i] = math.MaxInt64
	}

	for len(cellsToVisit) > 0 {
		// get the cell with the minimum total risk
		cellIndex := findMinimumRiskIndex(cellsToVisit)
		currentCell := cellsToVisit[cellIndex]
		newCellsToVisit := cellsToVisit[:cellIndex]
		newCellsToVisit = append(newCellsToVisit, cellsToVisit[cellIndex+1:]...)
		cellsToVisit = newCellsToVisit

		if currentCell.X == targetX && currentCell.Y == targetY {
			return currentCell.TotalRisk
		}

		// visit neighbors
		for _, neighborIndexes := range cellNeighbors {
			neighborX := currentCell.X + neighborIndexes[0]
			neighborY := currentCell.Y + neighborIndexes[1]
			neighborRisk := grid[neighborY][neighborX]
			neighborId := neighborY*gridWidth + neighborX

			// filter out grid limits
			if neighborRisk == math.MaxInt64 {
				continue
			}

			// check for a shorter path
			if currentCell.TotalRisk+neighborRisk < cellsTotalRisks[neighborId] {
				totalRisk := currentCell.TotalRisk + neighborRisk
				cellsTotalRisks[neighborId] = totalRisk
				neighbor := Cell{
					Id:        neighborY*gridWidth + neighborX,
					X:         neighborX,
					Y:         neighborY,
					TotalRisk: totalRisk,
				}
				cellsToVisit = append(cellsToVisit, neighbor)
			}
		}
	}
	return 0
}

func findMinimumRiskIndex(cells []Cell) int {
	minimumCellIndex := 0
	minimumRisk := cells[0].TotalRisk
	for index, cell := range cells {
		if cell.TotalRisk < minimumRisk {
			minimumRisk = cell.TotalRisk
			minimumCellIndex = index
		}
	}
	return minimumCellIndex
}

// Parse the grid and add a padding line/column
func parseBiggerInput(input []string) [][]int {
	repeatCount := 5
	innerGridHeight := len(input)
	innerGridWidth := len(input[0])
	gridHeight := (repeatCount * len(input)) + 2
	gridWidth := (repeatCount * len(input[0])) + 2

	grid := make([][]int, 0, gridHeight)

	emptyLine := make([]int, gridWidth)
	for i := range emptyLine {
		emptyLine[i] = math.MaxInt64
	}
	grid = append(grid, emptyLine)

	// first line of grids
	for _, lineStr := range input {
		line := make([]int, gridWidth)
		line[0] = math.MaxInt64

		for i, columnStr := range lineStr {
			number, _ := strconv.Atoi(string(columnStr))
			line[i+1] = number
		}

		// repeat the line to the right
		for gridRepeat := 1; gridRepeat < repeatCount; gridRepeat++ {
			for columnIndex := 0; columnIndex < innerGridWidth; columnIndex++ {
				risk := line[1+(gridRepeat-1)*innerGridWidth+columnIndex] + 1
				if risk > 9 {
					risk = 1
				}
				line[1+gridRepeat*innerGridWidth+columnIndex] = risk
			}
		}

		line[len(line)-1] = math.MaxInt64
		grid = append(grid, line)
	}

	// add remaining lines
	for gridRepeat := 1; gridRepeat < repeatCount; gridRepeat++ {
		for rowIndex := 0; rowIndex < innerGridHeight; rowIndex++ {
			line := make([]int, gridWidth)
			line[0] = math.MaxInt64
			for columnIndex := 1; columnIndex < gridWidth; columnIndex++ {
				risk := grid[1+(gridRepeat-1)*innerGridHeight+rowIndex][columnIndex] + 1
				if risk > 9 {
					risk = 1
				}
				line[columnIndex] = risk
			}
			line[len(line)-1] = math.MaxInt64
			grid = append(grid, line)
		}
	}

	grid = append(grid, emptyLine)
	return grid
}
