package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day7"
)

func init() {
	day := &cobra.Command{
		Use:   "7",
		Short: "Problems for Day 7",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day7.A(challenge.FromFile()))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day7.B(challenge.FromFile()))
		},
	})
	rootCmd.AddCommand(day)
}
