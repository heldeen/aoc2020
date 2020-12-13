package day12

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 441
func A(challenge *challenge.Input) int {

	s := ship{direction: 90}

	for l := range challenge.Lines() {
		s.navigate(l)
	}

	return util.Abs(s.x) + util.Abs(s.y)
}

type ship struct {
	direction, x, y int
}

func (s *ship) navigate(instruction string) {
	action := string(instruction[0])
	value := util.MustAtoI(instruction[1:])

	switch action {
	case "N":
		s.y += value
	case "S":
		s.y -= value
	case "E":
		s.x += value
	case "W":
		s.x -= value
	case "L":
		s.direction = s.direction - value
		if s.direction >= 360 {
			s.direction -= 360
		}
		if s.direction < 0 {
			s.direction += 360
		}
	case "R":
		s.direction = s.direction + value
		if s.direction >= 360 {
			s.direction -= 360
		}
	case "F":
		switch s.direction {
		case 0:
			s.y += value
		case 90:
			s.x += value
		case 180:
			s.y -= value
		case 270:
			s.x -= value
		}
	}
}
