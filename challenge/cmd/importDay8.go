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
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Parent().Flag("input")
			fmt.Printf("Answer: %d\n", day8.A(challenge.FromFileP(flag.Value.String())))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(cmd *cobra.Command, _ []string) {
			flag := cmd.Parent().Flag("input")
			fmt.Printf("Answer: %d\n", day8.B(challenge.FromFileP(flag.Value.String())))
		},
	})

	flags := day.Flags()

	flags.StringP("input", "i", "./challenge/day8/input.txt", "Input File to read")

	rootCmd.AddCommand(day)
}
