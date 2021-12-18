package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
target area: x=20..30, y=-10..-5
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 45, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 112, part2(input))
}

func Test_reverseSumToN(t *testing.T) {
	v, ok := reverseSumToN(10)
	require.Equal(t, 4, v)
	require.True(t, ok)

	v, ok = reverseSumToN(15)
	require.Equal(t, 5, v)
	require.True(t, ok)

	v, ok = reverseSumToN(11)
	require.Equal(t, 4, v)
	require.False(t, ok)
}
