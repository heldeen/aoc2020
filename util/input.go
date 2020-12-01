package util

import (
	"strconv"
)

func LinesToIntSlice(input <-chan string) (lines []int) {

	for l := range input {
		n, _ := strconv.Atoi(l)
		lines = append(lines, n)
	}
	return
}
