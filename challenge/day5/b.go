package day5

import (
	"fmt"
	"sort"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
	seatIds := calcSeatIds(challenge.Lines())

	sort.Ints(seatIds)

	for i := 1; i < len(seatIds); i++ {
		if seatIds[i]-seatIds[i-1] == 2 {
			fmt.Printf("found seatId: %d", seatIds[i]-1)
		}
	}
	return 0
}
