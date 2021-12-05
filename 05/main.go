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

	fmt.Println("Part 1:", countPart1Overlaps(input, 1000))
	fmt.Println("Part 2:", countPart2Overlaps(input, 1000))
}

// part 1
func countPart1Overlaps(segments []string, gridSize int) int {
	// for each line, extract segments
	// draw in the matrix
	grid := make([]int, gridSize*gridSize)
	for _, segment := range segments {
		parts := strings.Split(segment, " ")
		startCoordinatesStr := strings.Split(parts[0], ",")
		// omit the middle arrow
		endCoordinatesStr := strings.Split(parts[2], ",")
		x1, _ := strconv.Atoi(startCoordinatesStr[0])
		y1, _ := strconv.Atoi(startCoordinatesStr[1])
		x2, _ := strconv.Atoi(endCoordinatesStr[0])
		y2, _ := strconv.Atoi(endCoordinatesStr[1])

		// only draw horizontal or vertical lines
		if x1 == x2 {
			// swap values to ensure an ascending order
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				grid[y*gridSize+x1]++
			}
		} else if y1 == y2 {
			// swap values to ensure an ascending order
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				grid[y1*gridSize+x]++
			}
		}
	}

	// count overlaps
	overlapsCount := 0
	for _, n := range grid {
		if n >= 2 {
			overlapsCount++
		}
	}
	return overlapsCount
}

// part 2
func countPart2Overlaps(segments []string, gridSize int) int {
	// for each line, extract segments
	// draw in the matrix
	grid := make([]int, gridSize*gridSize)
	for _, segment := range segments {
		parts := strings.Split(segment, " ")
		startCoordinatesStr := strings.Split(parts[0], ",")
		// omit the middle arrow
		endCoordinatesStr := strings.Split(parts[2], ",")
		x1, _ := strconv.Atoi(startCoordinatesStr[0])
		y1, _ := strconv.Atoi(startCoordinatesStr[1])
		x2, _ := strconv.Atoi(endCoordinatesStr[0])
		y2, _ := strconv.Atoi(endCoordinatesStr[1])

		switch true {
		case x1 == x2:
			// swap values to ensure an ascending order
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				grid[y*gridSize+x1]++
			}
		case y1 == y2:
			// swap values to ensure an ascending order
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				grid[y1*gridSize+x]++
			}
		default:
			// swap values to ensure an ascending order
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			// diagonal lines are at 45°
			if y1 < y2 {
				for i := 0; i <= x2-x1; i++ {
					grid[(y1+i)*gridSize+(x1+i)]++
				}
			} else {
				for i := 0; i <= x2-x1; i++ {
					grid[(y1-i)*gridSize+(x1+i)]++
				}
			}
		}
	}

	// count overlaps
	overlapsCount := 0
	for _, n := range grid {
		if n >= 2 {
			overlapsCount++
		}
	}
	return overlapsCount
}
