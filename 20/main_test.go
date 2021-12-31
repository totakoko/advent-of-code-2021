package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 35, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 3351, part2(input))
}

func Test_parseInput(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	_, grid := parseInput(input, 10)
	F := 0
	T := 1
	expectedGrid := []int{
		F, F, F, F, F, F, F, F, F, F,
		F, F, F, F, F, F, F, F, F, F,
		F, F, T, F, F, T, F, F, F, F,
		F, F, T, F, F, F, F, F, F, F,
		F, F, T, T, F, F, T, F, F, F,
		F, F, F, F, T, F, F, F, F, F,
		F, F, F, F, T, T, T, F, F, F,
		F, F, F, F, F, F, F, F, F, F,
		F, F, F, F, F, F, F, F, F, F,
		F, F, F, F, F, F, F, F, F, F,
	}
	require.Equal(t, expectedGrid, grid.Pixels)
}
