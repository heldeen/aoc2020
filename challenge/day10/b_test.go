package day10

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

func TestB1(t *testing.T) {

	input := challenge.FromLiteral(input1)
	result := B(input)

	require.Equal(t, 8, result)

}

func TestB2(t *testing.T) {

	input := challenge.FromLiteral(input2)
	result := B(input)

	require.Equal(t, 19208, result)

}
