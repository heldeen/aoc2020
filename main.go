package main

import (
	"github.com/heldeen/aoc2020/challenge/day1"
	"os"

	"github.com/heldeen/aoc2020/challenge/cmd"
)

func main() {
	rootCommand := cmd.NewRootCommand()

	day1.AddCommandsTo(rootCommand)
	//dayX.AddCommandsTo(rootCommand)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
