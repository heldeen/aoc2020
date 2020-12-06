package util

import (
	"strconv"
	"strings"
)

func LinesToStrLine(input <-chan string) string {

	for l := range input {
		return l
	}
	return ""
}

func LinesToStrSlice(input <-chan string) (lines []string) {

	for l := range input {
		lines = append(lines, l)
	}
	return
}

func LinesToIntSlice(input <-chan string) (lines []int) {

	for l := range input {
		n, _ := strconv.Atoi(l)
		lines = append(lines, n)
	}
	return
}

func LinesToSingleCsvToStrSlice(input <-chan string) (csv []string) {
	return strings.Split(LinesToStrLine(input), ",")
}
