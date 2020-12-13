package day12

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

//Answer: 40014
func B(challenge *challenge.Input) int {
	s := waypointShip{waypointY: 1, waypointX: 10}

	for l := range challenge.Lines() {
		s.navigate(l)
	}

	return util.Abs(s.x) + util.Abs(s.y)
}

type waypointShip struct {
	x, y, waypointX, waypointY int
}

func swapSign(a int) int {
	if a > 0 {
		return -1 * a
	}
	return util.Abs(a)
}

func (s *waypointShip) navigate(instruction string) {
	action := string(instruction[0])
	value := util.MustAtoI(instruction[1:])

	switch action {
	case "N":
		s.waypointY += value
	case "S":
		s.waypointY -= value
	case "E":
		s.waypointX += value
	case "W":
		s.waypointX -= value
	case "L":
		switch value {

		case 90:
			wpX := swapSign(s.waypointY)
			wpY := s.waypointX

			s.waypointY = wpY
			s.waypointX = wpX

		case 180:
			s.waypointX = swapSign(s.waypointX)
			s.waypointY = swapSign(s.waypointY)
		case 270:
			wpX := s.waypointY
			wpY := swapSign(s.waypointX)

			s.waypointY = wpY
			s.waypointX = wpX
		}
	case "R":
		switch value {

		case 90:
			wpX := s.waypointY
			wpY := swapSign(s.waypointX)

			s.waypointY = wpY
			s.waypointX = wpX

		case 180:
			s.waypointX = swapSign(s.waypointX)
			s.waypointY = swapSign(s.waypointY)
		case 270:
			wpX := swapSign(s.waypointY)
			wpY := s.waypointX

			s.waypointY = wpY
			s.waypointX = wpX
		}
	case "F":
		s.x = s.x + s.waypointX*value
		s.y = s.y + s.waypointY*value
	}
}
