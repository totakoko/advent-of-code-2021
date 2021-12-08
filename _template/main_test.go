package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Result(t *testing.T) {
	input := strings.Split(`1 2 3`, "\n")
	require.Equal(t, 0, part1(input))
	require.Equal(t, 0, part2(input))
}
