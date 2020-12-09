package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day8"
)

func init() {
	day := &cobra.Command{
		Use:   "8",
		Short: "Problems for Day 8",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day8.A(challenge.FromFile()))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day8.B(challenge.FromFile()))
		},
	})
	rootCmd.AddCommand(day)
}
