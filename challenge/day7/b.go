package day7

import (
	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 34988
func B(challenge *challenge.Input) int {

	rules := loadRules(challenge.Lines())

	return countBags(rules, "shiny gold")
}

func countBags(rules rules, root string) (total int) {
	b := rules[root]
	for k, v := range b {
		total += v * (1 + countBags(rules, k))
	}
	return
}
