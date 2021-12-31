package main

import (
	"fmt"
	"io/ioutil"
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
	algorithm, grid := parseInput(input, 120)

	applyAlgorithm(algorithm, &grid)
	applyAlgorithm(algorithm, &grid)

	return grid.CountLitPixels()
}

func part2(input []string) int {
	algorithm, grid := parseInput(input, 300)

	for i := 0; i < 50; i++ {
		applyAlgorithm(algorithm, &grid)
	}

	return grid.CountLitPixels()
}

func parseInput(input []string, gridSize int) (string, Grid) {
	algorithm := input[0]
	imageSize := len(input[2])
	beginX := (gridSize - imageSize) / 2
	beginY := beginX * gridSize
	litValue := 1
	pixels := make([]int, gridSize*gridSize)
	for l, line := range input[2:] {
		for c, col := range line {
			if col == '#' {
				pixels[beginY+gridSize*l+beginX+c] = litValue
			}
		}
	}

	return algorithm, Grid{
		Pixels:           pixels,
		GridSize:         gridSize,
		ImageSize:        imageSize,
		LitValue:         litValue,
		EvolutionCounter: 0,
	}
}

type Grid struct {
	Pixels           []int
	GridSize         int
	ImageSize        int
	LitValue         int
	EvolutionCounter int
}

func (grid Grid) At(x int, y int) *int {
	return &grid.Pixels[grid.GridSize*y+x]
}

func (grid Grid) IsLit(x int, y int) bool {
	return grid.Pixels[grid.GridSize*y+x]&grid.LitValue > 0
}

func (grid Grid) Print() {
	// fmt.Print("\033[H\033[2J") // clear the terminal
	for y := 0; y < grid.GridSize; y++ {
		for x := 0; x < grid.GridSize; x++ {
			char := "."
			if grid.IsLit(x, y) {
				char = "#"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func (grid Grid) CountLitPixels() int {
	counter := 0

	// borders + padding are invalid, so we only check inner pixels
	padding := grid.EvolutionCounter
	beginIndex := (grid.GridSize-grid.ImageSize)/2 - padding
	for y := beginIndex; y < beginIndex+grid.ImageSize+2*padding; y++ {
		for x := beginIndex; x < beginIndex+grid.ImageSize+2*padding; x++ {
			if grid.IsLit(x, y) {
				counter++
			}
		}
	}

	return counter
}

func getNeighbors(x int, y int) [][2]int {
	return [][2]int{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

func applyAlgorithm(algo string, grid *Grid) {
	grid.EvolutionCounter++
	nextLitValue := grid.LitValue << 1 // limit at 63 steps
	for y := 1; y < grid.GridSize-1; y++ {
		for x := 1; x < grid.GridSize-1; x++ {
			binaryStr := ""
			for _, neighbor := range getNeighbors(x, y) {
				if grid.IsLit(neighbor[0], neighbor[1]) {
					binaryStr += "1"
				} else {
					binaryStr += "0"
				}
			}
			index, err := strconv.ParseInt(binaryStr, 2, 64)
			if err != nil {
				panic(err)
			}
			if algo[index] == '#' {
				*grid.At(x, y) |= nextLitValue
			}
		}
	}
	grid.LitValue = nextLitValue
}
