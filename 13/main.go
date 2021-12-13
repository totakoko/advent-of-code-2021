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
	dots, foldInstructions := parseInput(input)

	// apply the first fold only
	visibleDots := foldPaper(dots, foldInstructions[0])

	// count the visible dots
	return len(visibleDots)
}

func part2(input []string) int {
	dots, foldInstructions := parseInput(input)

	// apply the first fold only
	for _, foldInstruction := range foldInstructions {
		dots = foldPaper(dots, foldInstruction)
	}

	// manual recognition needed here
	printPaper(dots)
	return 0
}

type FoldInstruction struct {
	Axis  string
	Index int
}

func parseInput(input []string) ([][2]int, []FoldInstruction) {
	foldInstructionsIndex := 0
	dots := [][2]int{}
	for index, dotLine := range input {
		if dotLine == "" {
			foldInstructionsIndex = index + 1
			break
		}
		parts := strings.Split(dotLine, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		dots = append(dots, [2]int{x, y})
	}

	foldInstructions := []FoldInstruction{}
	for _, foldLine := range input[foldInstructionsIndex:] {
		parts := strings.Split(foldLine, "=")
		axis := string(parts[0][len(parts[0])-1])
		foldIndex, _ := strconv.Atoi(parts[1])
		foldInstructions = append(foldInstructions, FoldInstruction{
			Axis:  axis,
			Index: foldIndex,
		})
	}

	return dots, foldInstructions
}

func foldPaper(dots [][2]int, foldInstruction FoldInstruction) [][2]int {
	visibleDots := map[string][2]int{}
	foldIndex := foldInstruction.Index
	switch foldInstruction.Axis {
	case "x":
		// horizontal split
		for _, dot := range dots {
			if dot[0] > foldIndex {
				dot[0] = foldIndex - (dot[0] - foldIndex)
			}
			visibleDots[strconv.Itoa(dot[0])+","+strconv.Itoa(dot[1])] = dot
		}
	case "y":
		// vertical split
		for _, dot := range dots {
			if dot[1] > foldIndex {
				dot[1] = foldIndex - (dot[1] - foldIndex)
			}
			visibleDots[strconv.Itoa(dot[0])+","+strconv.Itoa(dot[1])] = dot
		}
	}

	visibleDotsSlice := make([][2]int, 0, len(visibleDots))
	for _, dot := range visibleDots {
		visibleDotsSlice = append(visibleDotsSlice, dot)
	}
	return visibleDotsSlice
}

func printPaper(dots [][2]int) {
	maxX := 0
	maxY := 0
	for _, dot := range dots {
		if dot[0] > maxX {
			maxX = dot[0] + 1
		}
		if dot[1] > maxY {
			maxY = dot[1] + 1
		}
	}

	grid := make([]bool, (maxX+1)*(maxY+1))
	for _, dot := range dots {
		grid[dot[1]*maxX+dot[0]] = true
	}

	for index, cell := range grid {
		if index%maxX == 0 {
			fmt.Println()
		}
		if cell {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
