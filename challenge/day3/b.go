package day3

import (
	"fmt"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 1260601650
func b(challenge *challenge.Input) int {
	h := loadHill(challenge.Lines())

	return h.traverse(1, 1) * h.traverse(3, 1) * h.traverse(5, 1) * h.traverse(7, 1) * h.traverse(1, 2)
}
