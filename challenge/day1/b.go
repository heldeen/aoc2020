package day1

import (
	"fmt"
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 193598720
func b(challenge *challenge.Input) int {

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
