package day10

import (
	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 193434623148032
func B(challenge *challenge.Input) int {

	adapters := getAdapters(challenge.Lines())

	pathsToNumber := map[int]int{
		0: 1,
	}

	for i := 0; i < len(adapters); i++ {
		current := adapters[i]
		for j := i + 1; j < len(adapters) && current+3 >= adapters[j]; j++ {
			next := adapters[j]
			pathsToNumber[next] += pathsToNumber[current]
		}
	}
	last := adapters[len(adapters)-1]

	return pathsToNumber[last]
}
