package day13

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

const input1 = `939
7,13,x,x,59,x,31,19`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(input1)

	result := A(input)

	require.Equal(t, 295, result)
}
