package day9

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 138879426
func A(challenge *challenge.Input) int {
	preambleLen := 25

	preamble := util.LinesToIntSlice(challenge.Lines())

	return xmas1(preambleLen, preamble)
}

func xmas1(preambleLen int, preamble []int) int {
	for i := preambleLen + 1; i < len(preamble); i++ {
		next := preamble[i]

		if xmas1a(next, preamble[i-preambleLen:i]) {
			preamble = append(preamble[1:], next)
			continue
		}

		return next
	}

	return 0
}

func xmas1a(next int, preamble []int) bool {
	for _, j := range preamble {
		for _, l := range preamble {
			if j == l {
				continue
			} else if j+l == next {
				return true
			}
		}
	}
	return false
}
