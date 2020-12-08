package day7

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/heldeen/aoc2020/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

//Answer: 177
func a(challenge *challenge.Input) int {
	rules := loadRules(challenge.Lines())

	masterList := util.NewSet()

	list := util.NewSet()
	list.Add("shiny gold")

	for list.Len() > 0 {
		list2 := util.NewSet()
		for _, l := range list.List() {
			for k, v := range rules {
				if _, found := v[l]; found {
					list2.Add(k)
				}
			}
		}

		masterList.AddAll(list2.List())

		list = list2
		list2 = nil
	}
	return masterList.Len()
}

type rules map[string]map[string]int

func loadRules(in <-chan string) rules {
	rules := make(map[string]map[string]int)

	for l := range in {
		bag, contains := parseLine(l)

		r, ok := rules[bag]

		if !ok {
			r = make(map[string]int)
		}
		for _, cr := range contains {
			count := util.MustAtoI(cr[0:1])
			b := cr[2:]
			br, ok2 := rules[b]
			if !ok2 {
				br = make(map[string]int)
				rules[b] = br
			}
			r[b] = count
		}
		rules[bag] = r
	}
	return rules
}

func parseLine(l string) (bag string, contains []string) {
	b := strings.Split(l, " bags contain ")

	bag = b[0]

	if strings.EqualFold("no other bags.", b[1]) {
		return
	}

	for _, c := range strings.Split(b[1][0:len(b[1])-1], ", ") {
		contains = append(contains, strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(c, "bags", ""), "bag", "")))
	}
	return
}
