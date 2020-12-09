package day3

import (
	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 1260601650
func B(challenge *challenge.Input) int {
	h := loadHill(challenge.Lines())

	return h.traverse(1, 1) * h.traverse(3, 1) * h.traverse(5, 1) * h.traverse(7, 1) * h.traverse(1, 2)
}
