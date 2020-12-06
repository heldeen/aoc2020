package day5

import (
	"fmt"
	"sort"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

//Answer: 919
func a(challenge *challenge.Input) int {

	seatIds := calcSeatIds(challenge.Lines())

	sort.Ints(seatIds)

	return seatIds[len(seatIds)-1]
}

const maxRow = 128
const maxColumn = 8

func seatId(row, column int) int {
	return row*8 + column
}

func calcSeatIds(in <-chan string) (seatIds []int) {
	for l := range in {
		id := calcSeatId(l)

		seatIds = append(seatIds, id)
	}
	return seatIds
}

func calcSeatId(l string) int {

	row := calcRow(l)

	column := calcColumn(l)

	id := seatId(row, column)
	return id
}

func calcColumn(l string) int {
	columns := make([]int, maxColumn)

	for i := 0; i < len(columns); i++ {
		columns[i] = i
	}
	for _, c := range l[7:] {
		switch string(c) {
		case "L":
			columns = columns[:len(columns)/2]
		case "R":
			columns = columns[len(columns)/2:]
		}
	}
	column := columns[0]
	return column
}

func calcRow(l string) int {
	rows := make([]int, maxRow)
	for i := 0; i < len(rows); i++ {
		rows[i] = i
	}
	for _, c := range l[0:8] {
		switch string(c) {
		case "F":
			rows = rows[:len(rows)/2]
		case "B":
			rows = rows[len(rows)/2:]
		}
	}
	row := rows[0]
	return row
}
