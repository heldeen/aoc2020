package day2

import (
	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 562
func B(challenge *challenge.Input) int {
	var valid int

	for l := range challenge.Lines() {
		first, second, letter, pass := parseLine(l)

		firstLetter := string(pass[first-1])
		secondLetter := string(pass[second-1])

		if firstLetter != secondLetter && (firstLetter == letter || secondLetter == letter) {
			valid++
		}
	}

	return valid
}
