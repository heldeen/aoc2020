package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day1"
)

func init() {
	day := &cobra.Command{
		Use:   "1",
		Short: "Problems for Day 1",
	}

	a := &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day1.A(challenge.FromFileP(flag.Value.String())))
		},
	}
	flagsA := a.Flags()
	flagsA.StringP("input", "i", "./challenge/day1/input.txt", "Input File to read")

	day.AddCommand(a)

	b := &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Flag("input")
			fmt.Printf("Answer: %d\n", day1.B(challenge.FromFileP(flag.Value.String())))
		},
	}

	flagsB := b.Flags()
	flagsB.StringP("input", "i", "./challenge/day1/input.txt", "Input File to read")

	day.AddCommand(b)

	rootCmd.AddCommand(day)
}
