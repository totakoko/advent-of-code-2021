package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 1588, dumbPolymerization(input, 10))
	require.Equal(t, 1588, improvedPolymerization(input, 10))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	// require.Equal(t, 2188189693529, dumbPolymerization(input, 40)) // boom
	require.Equal(t, 2188189693529, improvedPolymerization(input, 40))
}
