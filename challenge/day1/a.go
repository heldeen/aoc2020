package day1

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 32064
func A(challenge *challenge.Input) int {

	lines := util.LinesToIntSlice(challenge.Lines())

	for _, i := range lines {
		for _, j := range lines {
			if i+j == 2020 {
				return i * j
			}
		}
	}

	return 0
}
