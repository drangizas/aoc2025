package day02

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	totalSum := 0

	for _, rangeStr := range strings.Split(input, ",") {
		separatorIdx := strings.Index(rangeStr, "-")
		from, _ := strconv.Atoi(rangeStr[:separatorIdx])
		to, _ := strconv.Atoi(strings.TrimSpace(rangeStr[separatorIdx+1:]))

		for i := from; i <= to; i++ {
			if isEqual(strconv.Itoa(i)) {
				totalSum += i
			}
		}
	}

	return strconv.Itoa(int(totalSum))
}

func isEqual(number string) bool {
	length := len(number)

	if length%2 != 0 {
		return false
	}

	return number[:length/2] == number[length/2:]
}
