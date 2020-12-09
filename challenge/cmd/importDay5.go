package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day5"
)

func init() {
	day := &cobra.Command{
		Use:   "5",
		Short: "Problems for Day 5",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day5.A(challenge.FromFileP(flag.Value.String())))
		},
	}

	a.Flags().StringP("input", "i", "./challenge/day5/input.txt", "Input File to read")

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day5.B(challenge.FromFileP(flag.Value.String())))
		},
	}

	b.Flags().StringP("input", "i", "./challenge/day5/input.txt", "Input File to read")

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
