package day6

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 3464
func b(challenge *challenge.Input) int {
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
