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

//Lines go slices based on whitespace grouping
func GroupedLines(input <-chan string) <-chan []string {

	groupedInput := make(chan []string)
	go func() {
		var lines []string
		for l := range input {
			if len(l) == 0 {
				groupedInput <- lines
				lines = nil
			} else {
				lines = append(lines, l)
			}
		}
		groupedInput <- lines
		close(groupedInput)
	}()
	return groupedInput
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
