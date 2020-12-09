package day6

import (
	"strings"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 3464
func B(challenge *challenge.Input) int {
	sum := 0

	for gl := range util.GroupedLines(challenge.Lines()) {
		groupSize := 0
		questions := make(map[string]int)
		for _, l := range gl {
			groupSize++
			for _, c := range strings.Split(l, "") {
				count, _ := questions[c]
				questions[c] = count + 1
			}
		}
		sum += countEveryoneAnsweredYes(questions, groupSize)
	}

	return sum
}

func countEveryoneAnsweredYes(questions map[string]int, groupSize int) (sum int) {
	for _, v := range questions {
		if v == groupSize {
			sum++
		}
	}
	return sum
}
