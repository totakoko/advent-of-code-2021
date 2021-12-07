package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Result(t *testing.T) {
	input := strings.Split(`16,1,2,0,4,2,7,1,2,14`, "\n")
	require.Equal(t, 37, part1(input[0]))
	require.Equal(t, 168, part2(input[0]))
}
