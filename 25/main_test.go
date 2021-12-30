package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 58, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 0, part2(input))
}
