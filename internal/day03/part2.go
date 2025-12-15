package day03

import (
	"strconv"
	"strings"
)

const MaxCount = 12

func Part2(input string) string {
	var totalMax int64

	for line := range strings.Lines(input) {
		max := findMaxDigits(strings.TrimSpace(line))
		totalMax += max
	}

	return strconv.FormatInt(totalMax, 10)
}

func findMaxDigits(line string) int64 {
	lineLen := len(line)

	var findings [MaxCount]int

	prevIndexFound := -1
	for cntRemaining := MaxCount; cntRemaining > 0; cntRemaining-- {
		for i := prevIndexFound + 1; i <= lineLen-cntRemaining; i++ {
			parsed := int(line[i] - '0')

			currFindingsIdx := MaxCount - cntRemaining
			if parsed > findings[currFindingsIdx] {
				findings[currFindingsIdx] = parsed
				prevIndexFound = i
			}
		}
	}

	var total int64
	for _, value := range findings {
		total = total*10 + int64(value)
	}

	return total
}
