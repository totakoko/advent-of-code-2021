package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CountIncreases(t *testing.T) {
	report := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	require.Equal(t, 7, countIncreases(report))
}

func Test_CountSlidingWindowIncreases(t *testing.T) {
	report := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	require.Equal(t, 5, countSlidingWindowIncreases(report))
}
