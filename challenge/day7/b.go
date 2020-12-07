package day7

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 34988
func b(challenge *challenge.Input) int {

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
