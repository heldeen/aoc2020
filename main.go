package main

import (
	"os"

	"github.com/heldeen/aoc2020/challenge/cmd"
)

func main() {
	rootCommand := cmd.NewRootCommand()

	//dayX.AddCommandsTo(rootCommand)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
