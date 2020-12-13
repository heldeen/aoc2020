package day13

import (
	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
	"sort"
	"strings"
)

//Answer: 2215
func A(challenge *challenge.Input) int {

	earliest := util.MustAtoI(<-challenge.Lines())

	busesX := strings.Split(<-challenge.Lines(), ",")

	var buses []int
	for _, b := range busesX {
		if b == "x" {
			continue
		}
		buses = append(buses, util.MustAtoI(b))
	}

	sort.Ints(buses)

	busTimes := make(map[int]int)
	for _, b := range buses {
		for i := 1; ; i++ {
			bt := b * i
			if bt >= earliest {
				busTimes[bt-earliest] = b
				break
			}
		}
	}

	smallest := buses[len(buses)-1]

	for k, _ := range busTimes {
		if k < smallest {
			smallest = k
		}
	}

	return busTimes[smallest] * smallest
}
