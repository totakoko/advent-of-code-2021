package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const smallGraph = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

const mediumGraph = `
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

const bigGraph = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

func Test_Part1(t *testing.T) {
	require.Equal(t, 10, part1(strings.Split(strings.TrimSpace(smallGraph), "\n")))
	require.Equal(t, 19, part1(strings.Split(strings.TrimSpace(mediumGraph), "\n")))
	require.Equal(t, 226, part1(strings.Split(strings.TrimSpace(bigGraph), "\n")))
}

func Test_Part2(t *testing.T) {
	require.Equal(t, 36, part2(strings.Split(strings.TrimSpace(smallGraph), "\n")))
	require.Equal(t, 103, part2(strings.Split(strings.TrimSpace(mediumGraph), "\n")))
	require.Equal(t, 3509, part2(strings.Split(strings.TrimSpace(bigGraph), "\n")))
}
