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
	targetArea := parseInput(input)
	// the trajectory always comes back at 0
	// in that case, the maximum possible velocityY is the distance from 0 to the bottom of the area (negative)
	// thus, the initial velocityY is -minY - 1 (-1 to get the previous velocity = the initial velocity)
	velocityY := -targetArea.MinY - 1
	maxY := intSumToN(velocityY)
	return maxY
}

func part2(input []string) int {
	return 0
}

type Area struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func (area *Area) Contains(x int, y int) bool {
	return area.MinX <= x && x <= area.MaxX && area.MinY <= y && y <= area.MaxY
}

func parseInput(input []string) Area {
	parts := strings.Split(input[0], ", y=")
	yParts := strings.Split(parts[1], "..")
	xParts := strings.Split(strings.Split(parts[0], "x=")[1], "..")
	minX, _ := strconv.Atoi(xParts[0])
	maxX, _ := strconv.Atoi(xParts[1])
	minY, _ := strconv.Atoi(yParts[0])
	maxY, _ := strconv.Atoi(yParts[1])
	return Area{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
	}
}

func intSumToN(n int) int {
	return n * (n + 1) / 2
}
