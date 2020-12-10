package day10

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heldeen/aoc2020/challenge"
)

var input1 = `16
10
15
5
1
11
7
19
6
12
4`

var input2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestA1(t *testing.T) {
	input := challenge.FromLiteral(input1)
	result := A(input)

	require.Equal(t, 35, result)
}

func TestA2(t *testing.T) {

	input := challenge.FromLiteral(input2)

	result := A(input)

	require.Equal(t, 220, result)
}
