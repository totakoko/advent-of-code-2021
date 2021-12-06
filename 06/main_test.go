package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Result(t *testing.T) {
	input := strings.Split(`3,4,3,1,2`, "\n")
	// part 1
	require.Equal(t, 5934, evolvePopulation(parseInput(input[0]), 80))

	// part 2
	// require.Equal(t, 26984457539, evolvePopulation(parseInput(input[0]), 256)) // boom !
	require.Equal(t, 5934, evolvePopulationWithModel(parseInput(input[0]), 80))
	require.Equal(t, 26984457539, evolvePopulationWithModel(parseInput(input[0]), 256))

	require.Equal(t, int64(5934), evolveBigPopulationWithModel(parseInput(input[0]), 80).Int64())
	require.Equal(t, int64(26984457539), evolveBigPopulationWithModel(parseInput(input[0]), 256).Int64())
}
