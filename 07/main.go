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

	fmt.Println("# Part 1", part1(input[0]))
	fmt.Println("# Part 2", part2(input[0]))
}

func part1(input string) int {
	crabPositions := parsePositions(input)

	minFuel := math.MaxInt64
	for i := 0; i < len(crabPositions); i++ {
		fuelSum := 0
		for _, crabPosition := range crabPositions {
			fuel := abs(i - crabPosition)
			fuelSum += fuel
		}
		if fuelSum < minFuel {
			minFuel = fuelSum
		}
	}
	return minFuel
}

func part2(input string) int {
	crabPositions := parsePositions(input)

	minFuel := math.MaxInt64
	for i := 0; i < len(crabPositions); i++ {
		fuelSum := 0
		for _, crabPosition := range crabPositions {
			distance := abs(i - crabPosition)
			fuelSum += distanceToFuel(distance)
		}
		if fuelSum < minFuel {
			minFuel = fuelSum
		}
	}
	return minFuel
}

func parsePositions(input string) []int {
	crabPositions := []int{}
	for _, numberStr := range strings.Split(input, ",") {
		number, _ := strconv.Atoi(numberStr)
		crabPositions = append(crabPositions, number)
	}
	return crabPositions
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distanceToFuel(distance int) int {
	fuel := 0
	for i := 1; i <= distance; i++ {
		fuel += i
	}
	return fuel
}
