package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day12"
)

func init() {
	day := &cobra.Command{
		Use:   "12",
		Short: "Problems for Day 12",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 12, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day12.A(challenge.FromFileP(flag.Value.String())))
		},
	}

	a.Flags().StringP("input", "i", "./challenge/day12/input.txt", "Input File to read")

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 12, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day12.B(challenge.FromFileP(flag.Value.String())))
		},
	}

	b.Flags().StringP("input", "i", "./challenge/day12/input.txt", "Input File to read")

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
