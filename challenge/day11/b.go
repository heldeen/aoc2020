package day11

import (
	"github.com/heldeen/aoc2020/challenge"
	"log"
)

func B(challenge *challenge.Input) int {
	return modelPeopleB(loadSeating(challenge.Lines()))
}

func (s seating) getAdjacentVisibleSeats(x, y int) (adjacents []seat) {
	xMax := len(s[y]) - 1
	yMax := len(s) - 1
	if y > 0 {
		if x > 0 { //top left
			i := x - 1
			j := y - 1
			seat := s.getSeat(i, j)
			for seat == FLOOR && i > 0 && j > 0 {
				i -= 1
				j -= 1
				seat = s.getSeat(i, j)
			}
			adjacents = append(adjacents, seat)
		}

		if x != xMax { //top right
			i := x + 1
			j := y - 1
			seat := s.getSeat(i, j)
			for seat == FLOOR && i < xMax && j > 0 {
				i += 1
				j -= 1
				seat = s.getSeat(i, j)
			}
			adjacents = append(adjacents, seat)
		}
		//top
		j := y - 1
		seat := s.getSeat(x, j)
		for seat == FLOOR && j > 0 {
			j -= 1
			seat = s.getSeat(x, j)
		}

		adjacents = append(adjacents, seat)
	}
	if x > 0 { //left
		i := x - 1
		seat := s.getSeat(i, y)
		for seat == FLOOR && i > 0 {
			i -= 1
			seat = s.getSeat(i, y)
		}
		adjacents = append(adjacents, seat)
	}
	if x < xMax { //right
		i := x + 1
		seat := s.getSeat(i, y)
		for seat == FLOOR && i < xMax {
			i += 1
			seat = s.getSeat(i, y)
		}
		adjacents = append(adjacents, seat)
	}
	if y < yMax {
		if x > 0 { //bottom left
			i := x - 1
			j := y + 1
			seat := s.getSeat(i, j)
			for seat == FLOOR && i > 0 && j < yMax {
				i -= 1
				j += 1
				seat = s.getSeat(i, j)
			}
			adjacents = append(adjacents, seat)
		}
		if x < xMax { //bottom right
			i := x + 1
			j := y + 1
			seat := s.getSeat(i, j)
			for seat == FLOOR && i < xMax && j < yMax {
				i += 1
				j += 1
				seat = s.getSeat(i, j)
			}
			adjacents = append(adjacents, seat)
		}
		//bottom
		j := y + 1
		seat := s.getSeat(x, j)
		for seat == FLOOR && j < yMax {
			j += 1
			seat = s.getSeat(x, j)
		}

		adjacents = append(adjacents, seat)
	}

	return
}

func modelPeopleB(seats seating) int {
	change := true

	log.Print(seats.String())

	for change {

		change = false

		seatsC := seats.copy()

		x, y := seats.getDimensions()

		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				s := seats.getSeat(i, j)
				if s == FLOOR {
					continue
				}
				adjacentOccupied := countAdjacentOccupied(seats.getAdjacentVisibleSeats(i, j))
				if adjacentOccupied >= 5 {
					setSeat := seatsC.setSeat(i, j, EMPTY)
					change = change || setSeat
				} else if adjacentOccupied == 0 {
					setSeat := seatsC.setSeat(i, j, TAKEN)
					change = change || setSeat
				}
			}
		}
		seats = seatsC
		log.Print(seats.String())
	}
	return seats.countOccupied()
}
