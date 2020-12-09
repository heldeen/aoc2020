package day4

import (
	"strings"

	"github.com/heldeen/aoc2020/challenge"
)

//Answer: 226
func A(challenge *challenge.Input) int {

	valid := 0
	passports := parseRawPassports(challenge.Lines())
	for _, p := range passports {
		if validatePassport(p) {
			valid++
		}
	}

	return valid
}

func parseRawPassports(in <-chan string) []string {
	var raw []string

	var ppRaw string
	for l := range in {
		if len(strings.TrimSpace(l)) == 0 {
			raw = append(raw, strings.TrimSpace(ppRaw))
			ppRaw = ""
			continue
		}
		ppRaw = ppRaw + " " + l

	}
	if len(ppRaw) > 0 {
		raw = append(raw, strings.TrimSpace(ppRaw))
	}
	return raw
}

func validatePassport(raw string) bool {
	fields := strings.Fields(raw)

	if len(fields) < 7 {
		return false
	}

	var byr, iyr, eyr, hgt, hcl, ecl, pid /*, cid */ bool

	for _, f := range fields {
		s := strings.Split(f, ":")
		switch s[0] {
		case "byr":
			byr = true
		case "iyr":
			iyr = true
		case "eyr":
			eyr = true
		case "hgt":
			hgt = true
		case "hcl":
			hcl = true
		case "ecl":
			ecl = true
		case "pid":
			pid = true
			//case "cid":
			//	cid = true
		}
	}

	return byr && iyr && eyr && hgt && hcl && ecl && pid
}
