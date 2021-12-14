package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	// fmt.Println("# Part 1 (dumb)", dumbPolymerization(input, 10))
	fmt.Println("# Part 1", improvedPolymerization(input, 10))
	fmt.Println("# Part 2", improvedPolymerization(input, 40))
}

func dumbPolymerization(input []string, nbIterations int) int {
	polymer, rules := parseInput(input)
	polymer = polymer.Iterate(rules, nbIterations)
	return polymer.GetStats()
}

type Polymer string

func (polymer *Polymer) Iterate(rules []InsertionRule, nbIterations int) Polymer {
	for iteration := 1; iteration <= nbIterations; iteration++ {
		newPolymer := polymer.Apply(rules)
		polymer = &newPolymer
	}

	return *polymer
}

func (polymer *Polymer) Apply(rules []InsertionRule) Polymer {
	polymerBytes := []byte(*polymer)

	for i := 0; i < len(polymerBytes)-1; i++ {
		searchPattern := polymerBytes[i : i+2]

		// try to apply the first matching rule for the pattern
		for _, rule := range rules {
			if bytes.Equal(rule.Pattern, searchPattern) {
				newPolymerBytes := make([]byte, len(polymerBytes)+1)
				copy(newPolymerBytes, polymerBytes[:i+1])
				newPolymerBytes[i+1] = rule.Value
				copy(newPolymerBytes[i+2:], polymerBytes[i+1:])
				polymerBytes = newPolymerBytes
				i += 1
				break
			}
		}
	}

	return Polymer(string(polymerBytes))
}

func (polymer *Polymer) GetStats() int {
	nbElementsByChar := map[byte]int{}

	for _, char := range []byte(*polymer) {
		nbElementsByChar[char]++
	}

	maxNbElements := 0
	minNbElements := 9999
	for _, nbElements := range nbElementsByChar {
		if nbElements > maxNbElements {
			maxNbElements = nbElements
		}
		if nbElements < minNbElements {
			minNbElements = nbElements
		}
	}

	return maxNbElements - minNbElements
}

type InsertionRule struct {
	Pattern []byte
	Value   byte
}

func parseInput(input []string) (Polymer, []InsertionRule) {
	polymerTemplate := Polymer(input[0])
	insertions := []InsertionRule{}
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		insertions = append(insertions, InsertionRule{
			Pattern: []byte(parts[0]),
			Value:   byte(parts[1][0]),
		})
	}

	return polymerTemplate, insertions
}
