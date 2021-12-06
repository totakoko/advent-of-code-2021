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

	fmt.Println("# Population after 80 days", evolvePopulation(parseInput(input[0]), 80))
	fmt.Println("# Population after 256 days", evolvePopulationWithModel(parseInput(input[0]), 256))
}

func parseInput(input string) []int {
	numbersStr := strings.Split(input, ",")
	population := make([]int, 0, len(numbersStr))
	for _, numberStr := range numbersStr {
		number, _ := strconv.Atoi((numberStr))
		population = append(population, number)
	}
	return population
}

func evolvePopulation(population []int, days int) int {
	for day := 0; day < days; day++ {
		populationSize := len(population)
		for p := 0; p < populationSize; p++ {
			switch population[p] {
			case 0:
				population[p] = 6
				population = append(population, 8)
			default:
				population[p]--
			}
		}
	}
	return len(population)
}

func evolvePopulationWithModel(population []int, days int) int {
	populationByState := make([]int, 9)
	for _, fishState := range population {
		populationByState[fishState]++
	}

	for day := 0; day < days; day++ {
		populationByState = append(populationByState[1:], populationByState[0])
		populationByState[6] += populationByState[8]
	}
	totalPopulation := 0
	for _, populationSize := range populationByState {
		totalPopulation += populationSize
	}
	return totalPopulation
}
