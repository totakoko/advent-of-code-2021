package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Submarine(t *testing.T) {
	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	submarine := &Submarine{}
	for _, command := range commands {
		submarine.Move(extractMovementFromText(command))
	}
	require.Equal(t, 150, submarine.x*submarine.y)
}

func Test_ComplexSubmarine(t *testing.T) {
	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	submarine := &ComplexSubmarine{}
	for _, command := range commands {
		submarine.Move(extractMovementFromText(command))
	}
	require.Equal(t, 900, submarine.x*submarine.y)
}
