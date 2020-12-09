package day6

import (
	"strings"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 6878
func A(challenge *challenge.Input) int {

	sum := 0

	for gl := range util.GroupedLines(challenge.Lines()) {

		questions := make(map[string]bool, 26)
		for _, l := range gl {
			for _, c := range strings.Split(l, "") {
				questions[c] = true
			}
		}
		sum += len(questions)
	}

	return sum
}
