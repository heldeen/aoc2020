package day3

import (
	"strconv"

	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 181
func A(challenge *challenge.Input) int {

	h := loadHill(challenge.Lines())

	return h.traverse(3, 1)
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
