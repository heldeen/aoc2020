package day4

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/heldeen/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

//Answer: 160
func b(challenge *challenge.Input) int {
	valid := 0
	passports := parseRawPassports(challenge.Lines())
	for _, p := range passports {
		v := advancedValidatePassport(p)
		log.Printf("[%s] - %v\n\n", p, v)
		if v {
			valid++
		}
	}

	return valid
}

func advancedValidatePassport(raw string) bool {
	fields := strings.Fields(raw)

	if len(fields) < 7 {
		return false
	}

	var byr, iyr, eyr, hgt, hcl, ecl, pid /*, cid */ bool

	for _, f := range fields {
		s := strings.Split(f, ":")
		switch s[0] {
		case "byr":
			y, err := strconv.Atoi(s[1])
			if err != nil {
				break
			}
			if y < 1920 || y > 2002 {
				break
			}

			byr = true
			log.Printf("[%s] - %v four digits; at least 1920 and at most 2002\n", f, byr)
		case "iyr":
			y, err := strconv.Atoi(s[1])
			if err != nil {
				break
			}
			if y < 2010 || y > 2020 {
				break
			}
			iyr = true
			log.Printf("[%s] - %v (Issue Year) - four digits; at least 2010 and at most 2020\n", f, iyr)
		case "eyr":
			y, err := strconv.Atoi(s[1])
			if err != nil {
				break
			}
			if y < 2020 || y > 2030 {
				break
			}
			eyr = true
			log.Printf("[%s] - %v (Expiration Year) - four digits; at least 2020 and at most 2030\n", f, eyr)
		case "hgt":
			hgtValid := true
			rawH := s[1]
			rawHLen := len(rawH)
			if rawHLen < 3 {
				break
			}

			system := rawH[rawHLen-2:]

			measurement, _ := strconv.Atoi(rawH[:rawHLen-2])

			switch system {
			case "in":
				if measurement < 59 || measurement > 76 {
					hgtValid = false
				}
			case "cm":
				if measurement < 150 || measurement > 193 {
					hgtValid = false
				}
			default:
				hgtValid = false
			}

			hgt = hgtValid
			log.Printf("[%s] - %v (Height) - a number followed by either cm (at least 150 and at most 193) or in (at least 59 and at most 76)\n", f, hgtValid)
		case "hcl":
			hclRegex := regexp.MustCompile("#[\\da-f]{6}")

			hcl = hclRegex.MatchString(s[1])

			log.Printf("[%s] - %v (Hair Color) - a # followed by exactly six characters 0-9 or a-f\n", f, hcl)
		case "ecl":

			if s[1] == "amb" ||
				s[1] == "blu" ||
				s[1] == "brn" ||
				s[1] == "gry" ||
				s[1] == "grn" ||
				s[1] == "hzl" ||
				s[1] == "oth" {

				ecl = true
			}
			log.Printf("[%s] - %v (Eye Color) - exactly one of: amb blu brn gry grn hzl oth\n", f, ecl)
		case "pid":
			pidRegex := regexp.MustCompile("^[\\d]{9}$")

			pid = pidRegex.MatchString(s[1])

			log.Printf("[%s] - %v a nine-digit number, including leading zeroes\n", f, pid)

			//case "cid":
			//	cid = true
		}
	}

	return byr && iyr && eyr && hgt && hcl && ecl && pid
}
