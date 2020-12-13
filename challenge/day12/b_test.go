package day12

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(input1)

	result := B(input)

	require.Equal(t, 286, result)
}
