package day8

import (
	"strconv"
	"strings"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

type instruction struct {
	ins string
	arg int
}

type gameconsole struct {
	accumulator, insPointer int
	memory                  []instruction
}

func (g *gameconsole) Run() int {
	instructionTracker := util.NewSet()
	for g.insPointer < len(g.memory) {
		if instructionTracker.Contains(strconv.Itoa(g.insPointer)) {
			return g.accumulator
		}
		instructionTracker.Add(strconv.Itoa(g.insPointer))
		g.processInstruction(g.memory[g.insPointer].ins, g.memory[g.insPointer].arg)
	}
	return g.accumulator
}

func (g *gameconsole) processInstruction(ins string, arg int) int {
	var incInsPointer int
	switch ins {
	case "acc":
		g.accumulator += arg
		incInsPointer = 1
	case "jmp":
		incInsPointer = arg
	case "nop":
		incInsPointer = 1
	}
	g.insPointer += incInsPointer
	return 0
}

//Answer: 1137
func A(challenge *challenge.Input) int {

	var mem []instruction
	for l := range challenge.Lines() {
		raw := strings.Fields(l)
		mem = append(mem, instruction{
			ins: raw[0],
			arg: util.MustAtoI(raw[1]),
		})
	}

	g := gameconsole{
		memory: mem,
	}

	return g.Run()
}
