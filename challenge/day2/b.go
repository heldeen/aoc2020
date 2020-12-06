package day2

import (
	"fmt"
	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 562
func b(challenge *challenge.Input) int {
	var valid int

	for l := range challenge.Lines() {
		first, second, letter, pass := parseLine(l)

		firstLetter := string(pass[first-1])
		secondLetter := string(pass[second-1])

		if firstLetter != secondLetter && (firstLetter == letter || secondLetter == letter) {
			valid++
		}
	}

	return valid
}
