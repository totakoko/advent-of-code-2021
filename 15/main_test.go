package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 40, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 315, part2(input))
}
