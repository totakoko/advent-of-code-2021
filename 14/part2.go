package main

import (
	"math"
)

type StringCounter map[string]int

func (counter StringCounter) Add(key string, value int) {
	counter[key] += value
}

type Counter map[byte]int

func (counter Counter) Add(key byte, value int) {
	counter[key] += value
}

func improvedPolymerization(input []string, nbIterations int) int {
	polymer, rules := parseInput(input)
	polymerization := NewImprovedPolymerization(string(polymer), rules)
	polymerization.Iterate(nbIterations)
	return polymerization.GetStats()
}

type ImprovedPolymerization struct {
	template           string
	rulesByPair        map[string]byte
	nbReferencesByPair StringCounter
	nbReferencesByChar Counter
}

func NewImprovedPolymerization(template string, rules []InsertionRule) *ImprovedPolymerization {
	rulesByPair := map[string]byte{}
	for _, rule := range rules {
		rulesByPair[string(rule.Pattern)] = rule.Value
	}

	nbReferencesByPair := StringCounter{}
	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i : i+2])
		value := nbReferencesByPair[pair]
		if value == 0 {
			nbReferencesByPair[pair] = 0
		}
		nbReferencesByPair[pair]++
	}
	nbReferencesByChar := Counter{}
	for _, char := range template {
		nbReferencesByChar[byte(char)]++
	}

	return &ImprovedPolymerization{
		template,
		rulesByPair,
		nbReferencesByPair,
		nbReferencesByChar,
	}
}

func (polymer *ImprovedPolymerization) Iterate(nbIterations int) {
	for iteration := 1; iteration <= nbIterations; iteration++ {
		// fmt.Println("Iteration", iteration)
		newNbReferencesByPair := StringCounter{}
		for pair, nbReferences := range polymer.nbReferencesByPair {
			charToInsert, ok := polymer.rulesByPair[pair]
			if ok {
				newNbReferencesByPair.Add(string(pair[0])+string(charToInsert), nbReferences)
				newNbReferencesByPair.Add(string(charToInsert)+string(pair[1]), nbReferences)
				polymer.nbReferencesByChar.Add(charToInsert, nbReferences)
			}
		}
		polymer.nbReferencesByPair = newNbReferencesByPair
	}
}

func (polymer *ImprovedPolymerization) GetStats() int {
	maxNbReferences := 0
	minNbReferences := math.MaxInt64
	for _, nbReferences := range polymer.nbReferencesByChar {
		if nbReferences > maxNbReferences {
			maxNbReferences = nbReferences
		}
		if nbReferences < minNbReferences {
			minNbReferences = nbReferences
		}
	}

	return maxNbReferences - minNbReferences
}
