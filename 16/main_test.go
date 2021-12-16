package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Conversion(t *testing.T) {
	require.Equal(t, "110100101111111000101000", convertToBinaryString("D2FE28"))

}

func Test_Part1(t *testing.T) {
	// require.Equal(t, 16, part1("D2FE28")) // test
	// require.Equal(t, 16, part1("38006F45291200")) // test
	// require.Equal(t, 16, part1("EE00D40C823060")) // test

	require.Equal(t, 16, part1("8A004A801A8002F478"))
	require.Equal(t, 12, part1("620080001611562C8802118E34"))
	require.Equal(t, 23, part1("C0015000016115A2E0802F182340"))
	require.Equal(t, 31, part1("A0016C880162017C3686B18A3D4780"))
}

// func Test_Part2(t *testing.T) {
// 	require.Equal(t, 16, part2("8A004A801A8002F478"))
// 	require.Equal(t, 12, part2("620080001611562C8802118E34"))
// 	require.Equal(t, 23, part2("C0015000016115A2E0802F182340"))
// 	require.Equal(t, 31, part2("A0016C880162017C3686B18A3D4780"))
// }
