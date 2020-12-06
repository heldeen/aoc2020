package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/heldeen/aoc2020/challenge/cmd"
	"github.com/heldeen/aoc2020/challenge/day1"
	"github.com/heldeen/aoc2020/challenge/day2"
	"github.com/heldeen/aoc2020/challenge/day3"
	"github.com/heldeen/aoc2020/challenge/day4"
)

func main() {
	date := time.Now().Format("20060102150405")
	logger, err := os.OpenFile(fmt.Sprintf("AoC.%s.log", date), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logger)

	rootCommand := cmd.NewRootCommand()

	day1.AddCommandsTo(rootCommand)
	day2.AddCommandsTo(rootCommand)
	day3.AddCommandsTo(rootCommand)
	day3.AddCommandsTo(rootCommand)
	day4.AddCommandsTo(rootCommand)
	//dayX.AddCommandsTo(rootCommand) // go run gen/problem.go X a

	if err := rootCommand.Execute(); err != nil {
		log.Fatalln(err)
	}
}
