package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func toBinary(number int) string {
	return fmt.Sprintf("%05b", number)
}

func Test_processDiagnosticReport(t *testing.T) {
	report := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	gamma, epsilon := processDiagnosticReport(report)
	require.Equal(t, "10110", toBinary(gamma))
	require.Equal(t, "01001", toBinary(epsilon))
	require.Equal(t, 198, computePowerUsage(report))

	require.Equal(t, "10111", toBinary(computeOxygenGeneratorRating(report)))
	require.Equal(t, "01010", toBinary(computeCO2ScubberRating(report)))
}
