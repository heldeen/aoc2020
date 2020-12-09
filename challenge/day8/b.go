package day8

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", B(challenge.FromFile()))
		},
	}
}

func (g *gameconsole) RunB() int {
	instructionTracker := util.NewSet()
	for g.insPointer < len(g.memory) {
		if instructionTracker.Contains(strconv.Itoa(g.insPointer)) {
			return -1
		}
		instructionTracker.Add(strconv.Itoa(g.insPointer))
		g.processInstruction(g.memory[g.insPointer].ins, g.memory[g.insPointer].arg)
	}
	return g.accumulator
}

//Answer: 1125
func B(challenge *challenge.Input) int {
	var mem []instruction
	for l := range challenge.Lines() {
		raw := strings.Fields(l)
		mem = append(mem, instruction{
			ins: raw[0],
			arg: util.MustAtoI(raw[1]),
		})
	}
	for i := 0; i < len(mem); i++ {

		mem2 := make([]instruction, len(mem))

		copy(mem2, mem)

		if strings.EqualFold("jmp", mem2[i].ins) {
			mem2[i].ins = "nop"
		} else {
			continue
		}

		g := gameconsole{
			memory: mem2,
		}
		result := g.RunB()
		if -1 == result {
			continue
		}
		return result
	}
	return 0
}
