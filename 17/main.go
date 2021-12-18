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
	// the trajectory always comes back at 0
	// in that case, the maximum possible velocityY is the distance from 0 to the bottom of the area (negative)
	// thus, the initial velocityY is -minY - 1 (-1 to get the previous velocity = the initial velocity)
	velocityY := -targetArea.MinY - 1
	maxY := intSumToN(velocityY)
	return maxY
}

func part2(input []string) int {
	targetArea := parseInput(input)
	velocities := map[string]bool{}
	minVelocityX, _ := reverseSumToN(targetArea.MinX)
	for vX := minVelocityX; vX <= targetArea.MaxX; vX++ {
		for vY := targetArea.MinY; vY <= -targetArea.MinY-1; vY++ {
			maxProbeX := intSumToN(vX)
			for nbSteps := 1; nbSteps <= 500; nbSteps++ {
				probeX := nbSteps*vX - intSumToN(nbSteps-1)
				if nbSteps > vX {
					probeX = maxProbeX
				}
				probeY := nbSteps*vY - intSumToN(nbSteps-1)
				if probeY < targetArea.MinY {
					break
				}
				if targetArea.Contains(probeX, probeY) {
					velocities[strconv.Itoa(vX)+":"+strconv.Itoa(vY)] = true
				}
			}
		}
	}

	return len(velocities)
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

func reverseSumToN(sum int) (int, bool) {
	delta := 1 - 4*(-2*sum) // b² - 4(a=1)(c=-2sum)
	sqrt := math.Sqrt(float64(delta))
	if sqrt > 0 {
		x1 := (-1 - sqrt) / 2
		x2 := (-1 + sqrt) / 2
		intResult := true
		if x1 != math.Trunc(x1) || x2 != math.Trunc(x2) {
			intResult = false
		}
		if x2 > x1 {
			return int(x2), intResult
		}
		return int(x1), intResult
	}
	if sqrt == 0 {
		// not good! float
		return -1 / 2, false
	}
	// sqrt < 0
	// not good!
	return 0, false
}
