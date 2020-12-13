package day13

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

func TestB(t *testing.T) {

	require.Equal(t, 3417, B(challenge.FromLiteral("foobar\n17,x,13,19")))
	require.Equal(t, 754018, B(challenge.FromLiteral("foobar\n67,7,59,61")))
	require.Equal(t, 779210, B(challenge.FromLiteral("foobar\n67,x,7,59,61")))
	require.Equal(t, 1202161486, B(challenge.FromLiteral("foobar\n1789,37,47,1889")))
	require.Equal(t, 1068781, B(challenge.FromLiteral("foobar\n7,13,x,x,59,x,31,19")))
}
