package day11

import (
	"github.com/heldeen/aoc2020/challenge"
	"log"
	"strings"
)

//Answer: 2183
func A(challenge *challenge.Input) int {
	return modelPeopleA(loadSeating(challenge.Lines()))
}

type seat string

const (
	EMPTY = seat("L")
	TAKEN = seat("#")
	FLOOR = seat(".")
)

type seating [][]seat

func (s seating) getSeat(x, y int) seat {
	return s[y][x]
}

func (s seating) setSeat(x, y int, seat seat) bool {
	v := s[y][x]
	if v == FLOOR {
		panic("THAT'S FLOOR")
	}

	s[y][x] = seat

	return !(seat == v)
}

func (s seating) copy() seating {
	var c [][]seat
	for _, row := range s {
		cRow := make([]seat, len(row))
		_ = copy(cRow, row)
		c = append(c, cRow)
	}
	return c
}

func (s seating) getDimensions() (x, y int) {
	return len(s[0]), len(s)
}

func (s seating) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for _, row := range s {
		for _, seat := range row {
			sb.WriteString(string(seat))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s seating) countOccupied() (occupied int) {
	for _, row := range s {
		for _, seat := range row {
			if seat == TAKEN {
				occupied++
			}
		}
	}
	return
}

func (s seating) getAdjacentSeats(x, y int) (adjacents []seat) {
	if y != 0 {
		if x != 0 {
			adjacents = append(adjacents, s.getSeat(x-1, y-1))
		}
		if x != len(s[y])-1 {
			adjacents = append(adjacents, s.getSeat(x+1, y-1))
		}
		adjacents = append(adjacents, s.getSeat(x, y-1))
	}
	if x != 0 {
		adjacents = append(adjacents, s.getSeat(x-1, y))
	}
	if x != len(s[y])-1 {
		adjacents = append(adjacents, s.getSeat(x+1, y))
	}
	if y != len(s)-1 {
		if x != 0 {
			adjacents = append(adjacents, s.getSeat(x-1, y+1))
		}
		if x != len(s[y])-1 {
			adjacents = append(adjacents, s.getSeat(x+1, y+1))
		}
		adjacents = append(adjacents, s.getSeat(x, y+1))
	}
	return
}

func loadSeating(in <-chan string) (seats seating) {
	for l := range in {
		var row []seat
		for _, s := range l {
			row = append(row, seat(s))
		}
		seats = append(seats, row)
	}
	return
}

func modelPeopleA(seats seating) int {
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
				adjacentOccupied := countAdjacentOccupied(seats.getAdjacentSeats(i, j))
				if adjacentOccupied >= 4 {
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

func countAdjacentOccupied(seats []seat) (count int) {
	for _, s := range seats {
		if s == TAKEN {
			count++
		}
	}
	return
}
