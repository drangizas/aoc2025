package day03

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	totalMax := 0

	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		max1st, max1stIdx := findMax(line[:len(line)-1])
		max2nd, _ := findMax(line[max1stIdx+1:])

		max := max1st*10 + max2nd
		totalMax += max
	}

	return strconv.Itoa(totalMax)
}

func findMax(partial string) (int, int) {
	max, maxIdx := 0, 0

	for i := 0; i < len(partial); i++ {
		number, _ := strconv.Atoi(partial[i : i+1])

		if number > max {
			max = number
			maxIdx = i
		}
	}

	return max, maxIdx
}
