package day2

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"regexp"
)

//Answer: 398
func A(challenge *challenge.Input) int {

	var valid int

	for l := range challenge.Lines() {

		start, end, letter, pass := parseLine(l)

		sm2 := regexp.MustCompile(letter)
		matches := sm2.FindAllStringSubmatchIndex(pass, -1)

		if start <= len(matches) && end >= len(matches) {
			valid++
		}
	}

	return valid
}

var r = regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): (\\w+)")

func parseLine(line string) (p1, p2 int, p3, pass string) {
	s := r.FindAllStringSubmatch(line, -1)

	return util.MustAtoI(s[0][1]), util.MustAtoI(s[0][2]), s[0][3], s[0][4]
}
