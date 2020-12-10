package day10

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"sort"
)

func A(challenge *challenge.Input) int {

	adapters := getAdapters(challenge.Lines())

	jolt1, jolt3 := 0, 0
	for i := 0; i < len(adapters)-1; i++ {
		switch adapters[i+1] - adapters[i] {
		case 1:
			jolt1++
		case 3:
			jolt3++
		}
	}

	return jolt1 * jolt3
}

func getAdapters(in <-chan string) []int {
	adapters := util.LinesToIntSlice(in)

	adapters = append([]int{0}, adapters...)

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters
}
