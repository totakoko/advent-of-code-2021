package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Day5(t *testing.T) {
	input := strings.Split(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, "\n")
	require.Equal(t, 5, countPart1Overlaps(input, 10))
	require.Equal(t, 12, countPart2Overlaps(input, 10))
}
