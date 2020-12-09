package day1

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 193598720
func B(challenge *challenge.Input) int {

	lines := util.LinesToIntSlice(challenge.Lines())

	for _, i := range lines {
		for _, j := range lines {
			for _, k := range lines {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}

	return 0
}
