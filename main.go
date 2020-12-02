package main

import (
	"os"

	"github.com/heldeen/aoc2020/challenge/cmd"
	"github.com/heldeen/aoc2020/challenge/day1"
	"github.com/heldeen/aoc2020/challenge/day2"
)

func main() {
	rootCommand := cmd.NewRootCommand()

	day1.AddCommandsTo(rootCommand)
	day2.AddCommandsTo(rootCommand)
	//dayX.AddCommandsTo(rootCommand)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
