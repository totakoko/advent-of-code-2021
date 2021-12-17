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

func part1(input []string) int {
	targetArea := parseInput(input)

	// find velocityX so that we can reach targetArea.MinX
	velocityX := 0
	probeX := 0
	for probeX < targetArea.MinX {
		velocityX++ // TODO handle negative
		probeX = intSumToN(velocityX)
	}

	maxVelocityY := 0
	for velocityY := 0; velocityY < 300; velocityY++ { // bruteforce !
		for sleepSteps := 0; sleepSteps < 300; sleepSteps++ {
			nbSteps := int(math.Abs(float64(velocityX))) + sleepSteps
			negativeY := intSumToN(nbSteps - 1)
			if targetArea.Contains(probeX, nbSteps*velocityY-negativeY) {
				if velocityY > maxVelocityY {
					maxVelocityY = velocityY
				}
			}
		}
	}

	return intSumToN(maxVelocityY)
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
