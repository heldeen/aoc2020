package day6

import (
	"fmt"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

//Answer: 6878
func a(challenge *challenge.Input) int {

	sum := 0

	questions := make(map[string]bool, 26)

	for l := range challenge.Lines() {
		if len(l) == 0 {
			sum += len(questions)
			questions = make(map[string]bool, 26)
		} else {
			for _, c := range l {
				questions[string(c)] = true
			}
		}
	}
	sum += len(questions)

	return sum
}
