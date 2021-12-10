package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 26397, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 288957, part2(input))
}
