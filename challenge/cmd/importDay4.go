package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day4"
)

func init() {
	day := &cobra.Command{
		Use:   "4",
		Short: "Problems for Day 4",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Parent().Flag("input")
			fmt.Printf("Answer: %d\n", day4.A(challenge.FromFileP(flag.Value.String())))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Parent().Flag("input")
			fmt.Printf("Answer: %d\n", day4.B(challenge.FromFileP(flag.Value.String())))
		},
	})
	flags := day.Flags()
	flags.StringP("input", "i", "./challenge/day4/input.txt", "Input File to read")
	rootCmd.AddCommand(day)
}
