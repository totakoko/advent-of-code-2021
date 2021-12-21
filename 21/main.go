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
	p1, p2 := parseInput(input)

	// start at index 0
	p1--
	p2--

	die := Die(0)
	score1 := 0
	score2 := 0

	nbRolls := 0
	for steps := 1; steps < 10000; steps++ {
		move := (die.Roll() + die.Roll() + die.Roll())
		p1 = (p1 + move) % 10
		nbRolls += 3
		score1 += p1 + 1
		if score1 >= 1000 {
			return score2 * nbRolls
		}
		p2 = (p2 + (die.Roll() + die.Roll() + die.Roll())) % 10
		nbRolls += 3
		score2 += p2 + 1
		if score2 >= 1000 {
			return score1 * nbRolls
		}
	}

	return 0
}

type Die int

func (die *Die) Roll() int {
	*die++
	if *die == 101 {
		*die = 1
	}
	return int(*die)
}

func part2(input []string) int {
	return 0
}

func parseInput(input []string) (int, int) {
	p1, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	p2, _ := strconv.Atoi(strings.Split(input[1], ": ")[1])
	return p1, p2
}
