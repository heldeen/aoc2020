package day12

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

const input1 = `F10
N3
F7
R90
F11`

func TestA(t *testing.T) {

	input := challenge.FromLiteral(input1)

	result := A(input)

	require.Equal(t, 25, result)
}
