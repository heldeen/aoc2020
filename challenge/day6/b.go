package day6

import (
	"fmt"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
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
	sum, groupSize := 0, 0

	questions := make(map[string]int, 26)

	for l := range challenge.Lines() {
		if len(l) == 0 {
			sum += countEveryoneAnsweredYes(questions, groupSize)
			questions = make(map[string]int, 26)
			groupSize = 0
		} else {
			groupSize++
			for _, c := range l {
				count, _ := questions[string(c)]
				questions[string(c)] = count + 1
			}
		}
	}
	sum += countEveryoneAnsweredYes(questions, groupSize)

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
