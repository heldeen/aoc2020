package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/heldeen/aoc2020/challenge/cmd"
)

func main() {
	date := time.Now().Format("20060102150405")
	logger, err := os.OpenFile(fmt.Sprintf("AoC.%s.log", date), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logger)

	//dayX.AddCommandsTo(rootCommand) // go run gen/problem.go X a

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
