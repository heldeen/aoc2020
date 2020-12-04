package main

import (
	"os"

	"github.com/heldeen/aoc2020/challenge/cmd"
	"github.com/heldeen/aoc2020/challenge/day1"
	"github.com/heldeen/aoc2020/challenge/day2"
	"github.com/heldeen/aoc2020/challenge/day3"
)

func main() {
	rootCommand := cmd.NewRootCommand()

	day1.AddCommandsTo(rootCommand)
	day2.AddCommandsTo(rootCommand)
	day3.AddCommandsTo(rootCommand)
	day3.AddCommandsTo(rootCommand)
	//dayX.AddCommandsTo(rootCommand) // go run gen/problem.go 3 a

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
