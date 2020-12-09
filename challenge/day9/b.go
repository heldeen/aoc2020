package day9

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"sort"
)

//Answer: 23761694
func B(challenge *challenge.Input) int {
	preambleLen := 25

	preamble := util.LinesToIntSlice(challenge.Lines())

	return xmas2(xmas1(preambleLen, preamble), preamble)
}

func xmas2(xmas1Val int, preamble []int) int {

	for i := 0; i < len(preamble); i++ {
		for j := 1; j < len(preamble) && j+i < len(preamble); j++ {
			contiguous := preamble[i : j+i]
			sum := xmas2a(contiguous)
			if sum == xmas1Val {
				return xmas2b(contiguous)
			} else if sum > xmas1Val {
				continue
			}
		}
	}
	return 0
}

func xmas2a(contiguous []int) (sum int) {
	for i := 0; i < len(contiguous); i++ {
		sum += contiguous[i]
	}
	return
}

func xmas2b(contiguous []int) int {
	sort.Ints(contiguous)
	return contiguous[0] + contiguous[len(contiguous)-1]
}
