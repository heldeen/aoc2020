package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/challenge/day9"
)

func init() {
	day := &cobra.Command{
		Use:   "9",
		Short: "Problems for Day 9",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day9.A(challenge.FromFile()))
		},
	})
	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", day9.B(challenge.FromFile()))
		},
	})

	flags := day.PersistentFlags()

	flags.StringP("input", "i", "./challenge/day9/input.txt", "Input File to read")
	_ = viper.BindPFlags(flags)
	rootCmd.AddCommand(day)
}
