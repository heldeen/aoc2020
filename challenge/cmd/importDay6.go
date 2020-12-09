package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day6"
)

func init() {
	day := &cobra.Command{
		Use:   "6",
		Short: "Problems for Day 6",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day6.A(challenge.FromFileP(flag.Value.String())))
		},
	}

	a.Flags().StringP("input", "i", "./challenge/day6/input.txt", "Input File to read")

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day6.B(challenge.FromFileP(flag.Value.String())))
		},
	}

	b.Flags().StringP("input", "i", "./challenge/day6/input.txt", "Input File to read")

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
