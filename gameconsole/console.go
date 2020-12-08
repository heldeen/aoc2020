package gameconsole

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/heldeen/aoc2020/util"
)

func NewGameConsole(instructions []instruction) *console {
	return &console{
		memory: instructions,
	}
}

func LoadProgram(in <-chan string) []instruction {
	var mem []instruction
	for l := range in {
		raw := strings.Fields(l)
		mem = append(mem, instruction{
			ins: raw[0],
			arg: util.MustAtoI(raw[1]),
		})
	}
	return mem
}

type instruction struct {
	ins string
	arg int
}

func (i instruction) String() string {
	return fmt.Sprintf("Ins[%s %+-4d]", i.ins, i.arg)
}

type console struct {
	accumulator, insPointer, executionCounter int
	memory                                    []instruction
}

func (g *console) Run() int {
	instructionTracker := util.NewSet()
	for g.insPointer < len(g.memory) {
		if instructionTracker.Contains(strconv.Itoa(g.insPointer)) {
			return g.accumulator
		}
		instructionTracker.Add(strconv.Itoa(g.insPointer))
		g.executionCounter++
		origInsPointer := g.insPointer
		i := g.memory[origInsPointer]
		log.Printf("Execution:%05d >> MEM[%05d]:%v ACCUMULATOR:%+-5d", g.executionCounter, origInsPointer, i, g.accumulator)
		deltaIP := g.processInstruction(i.ins, i.arg)
		log.Printf(" ------ UPDATES >> InstructionPointer:%+-5d MEM[%05d]:%v ACCUMULATOR:%+-5d", deltaIP, origInsPointer, g.memory[origInsPointer], g.accumulator)
	}
	return g.accumulator
}

func (g *console) processInstruction(ins string, arg int) int {
	var incInsPointer int
	var sb strings.Builder
	sb.WriteString(" >> PROCESSING ->> ")
	switch ins {
	case "acc":
		sb.WriteString(fmt.Sprintf("acc += %+d", arg))
		g.accumulator += arg
		incInsPointer = 1
	case "jmp":
		sb.WriteString(fmt.Sprintf("jmp = %+d", arg))
		incInsPointer = arg
	case "nop":
		sb.WriteString(fmt.Sprintf("jmp = %+d", arg))
		incInsPointer = 1
	}
	g.insPointer += incInsPointer
	log.Printf(sb.String())
	return incInsPointer
}
