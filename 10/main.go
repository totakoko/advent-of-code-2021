package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	fmt.Println("# Part 1", part1(input))
	fmt.Println("# Part 2", part2(input))
}

var matchingPair = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',

	// debug
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var illegalCharacterPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func part1(input []string) int {
	totalPoints := 0
	for _, line := range input {
		totalPoints += getLineScore(line)
	}
	return totalPoints
}

func getLineScore(line string) int {
	stack := []rune{}
	for _, char := range line {
		switch char {
		case '(', '[', '{', '<':
			stack = append(stack, char)
		case ')', ']', '}', '>':
			if matchingPair[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1] // pop the stack
			} else {
				fmt.Printf("Expected %s, but found %s instead.\n", string(matchingPair[stack[len(stack)-1]]), string(char))
				return illegalCharacterPoints[char]
			}
		}
	}
	return 0
}

func part2(input []string) int {
	scores := []int{}
	for _, line := range input {
		if score := getIncompleteLineScore(line); score != 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	middleScore := scores[len(scores)/2]
	return middleScore
}

var autocompleteCharacterPoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func getIncompleteLineScore(line string) int {
	stack := []rune{}
	for _, char := range line {
		switch char {
		case '(', '[', '{', '<':
			stack = append(stack, char)
		case ')', ']', '}', '>':
			if matchingPair[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1] // pop the stack
			} else {
				return 0
			}
		}
	}

	if len(stack) > 0 {
		// incomplete line, complete the stack
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + autocompleteCharacterPoints[matchingPair[stack[i]]]
		}
		return score
	}
	return 0
}
