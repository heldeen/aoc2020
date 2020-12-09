package day5

import (
	"sort"

	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 642
func B(challenge *challenge.Input) int {
	seatIds := calcSeatIds(challenge.Lines())

	sort.Ints(seatIds)

	for i := 1; i < len(seatIds); i++ {
		if seatIds[i]-seatIds[i-1] == 2 {
			return seatIds[i] - 1
		}
	}
	return 0
}
