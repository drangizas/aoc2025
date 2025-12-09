package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(input string) string {
	totalSum := 0

	for _, rangeStr := range strings.Split(input, ",") {
		separatorIdx := strings.Index(rangeStr, "-")
		from, _ := strconv.Atoi(rangeStr[:separatorIdx])
		to, _ := strconv.Atoi(strings.TrimSpace(rangeStr[separatorIdx+1:]))

		for i := from; i <= to; i++ {
			if containsSequence(strconv.Itoa(i)) {
				fmt.Printf("Invalid ID: %v\n", i)
				totalSum += i
			}
		}
	}

	return strconv.Itoa(int(totalSum))
}

func containsSequence(number string) bool {
	length := len(number)

	for currentWindow := 1; currentWindow <= length/2; currentWindow++ {

		if length%currentWindow != 0 {
			continue
		}

		isSequence := true
		for chunk := 0; chunk <= length/currentWindow-2; chunk++ {
			firstChunkFrom := chunk * currentWindow
			firstChunkTo := (chunk + 1) * currentWindow
			secondChunkTo := (chunk + 2) * currentWindow

			if number[firstChunkFrom:firstChunkTo] != number[firstChunkTo:secondChunkTo] {
				isSequence = false
			}
		}

		if isSequence {
			return true
		}
	}

	return false
}
