package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const inputText = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

func Test_Part1(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 4140, part1(input))
}

func Test_Part2(t *testing.T) {
	input := strings.Split(strings.TrimSpace(inputText), "\n")
	require.Equal(t, 3993, part2(input))
}

func Test_GetMagnitude(t *testing.T) {
	require.Equal(t, 143, ReadNode("[[1,2],[[3,4],5]]").GetMagnitude())
	require.Equal(t, 1384, ReadNode("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]").GetMagnitude())
	require.Equal(t, 445, ReadNode("[[[[1,1],[2,2]],[3,3]],[4,4]]").GetMagnitude())
	require.Equal(t, 791, ReadNode("[[[[3,0],[5,3]],[4,4]],[5,5]]").GetMagnitude())
	require.Equal(t, 1137, ReadNode("[[[[5,0],[7,4]],[5,5]],[6,6]]").GetMagnitude())
	require.Equal(t, 3488, ReadNode("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]").GetMagnitude())
}

func Test_Copy(t *testing.T) {
	nodeA := ReadNode("[[1,2],[[3,4],5]]")
	nodeB := nodeA.Copy()
	nodeA.Left.Left.Value = 5
	require.Equal(t, 1, nodeB.Left.Left.Value)
}

type Test struct {
	Input          string
	ExpectedResult string
}

func Test_Add(t *testing.T) {
	tests := []Test{
		{
			Input: `
[1,1]
[2,2]
[3,3]
[4,4]
		      `,
			ExpectedResult: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			Input: `
[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
		      `,
			ExpectedResult: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			Input: `
[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]
		      `,
			ExpectedResult: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			Input: `
[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]
      `,
			ExpectedResult: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			Input: `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
      `,
			ExpectedResult: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}

	for _, test := range tests {
		input := strings.Split(strings.TrimSpace(test.Input), "\n")
		require.Equal(t, test.ExpectedResult, AddAll(parseInput(input)).String())
	}
}

func Test_Reduce(t *testing.T) {
	tests := []Test{
		{
			Input:          "[[[[[9,8],1],2],3],4]",
			ExpectedResult: "[[[[0,9],2],3],4]",
		},
		{
			Input:          "[7,[6,[5,[4,[3,2]]]]]",
			ExpectedResult: "[7,[6,[5,[7,0]]]]",
		},
		{
			Input:          "[[6,[5,[4,[3,2]]]],1]",
			ExpectedResult: "[[6,[5,[7,0]]],3]",
		},
		{
			Input:          "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			ExpectedResult: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.ExpectedResult, ReadNode(test.Input).Reduce().String())
	}
}
