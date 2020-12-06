package day6

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
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
