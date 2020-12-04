package day3

import (
	"fmt"
	"strconv"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

type hill struct {
	grid     [][]bool
	y, realX int
}

func (h hill) treeAt(x, y int) bool {

	if y >= h.y {
		panic("y is too big! " + strconv.Itoa(y))
	}

	return h.grid[y][x%(h.realX)]
}

//Answer: 181
func a(challenge *challenge.Input) int {

	h := loadHill(challenge.Lines())

	return h.traverse(3, 1)
}

func (h hill) traverse(xInc, yInc int) int {
	trees, x, y := 0, 0, 0

	for y < h.y {
		if h.treeAt(x, y) {
			trees++
		}
		x += xInc
		y += yInc
	}

	return trees
}

func loadHill(in <-chan string) hill {
	var grid [][]bool
	var width int
	for l := range in {
		width = len(l)
		x := make([]bool, 0)
		for _, c := range l {
			x = append(x, string(c) == "#")
		}
		grid = append(grid, x)
	}

	return hill{
		grid:  grid,
		y:     len(grid),
		realX: width,
	}
}
