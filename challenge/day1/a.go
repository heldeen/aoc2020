package day1

import (
	"fmt"
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

//Answer: 32064
func a(challenge *challenge.Input) int {

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
