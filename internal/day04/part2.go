package day04

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	var count int

	done := false
	for !done {
		var recentLoopCnt int
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] != '@' {
					continue
				}

				var adjNeighborsCnt int

				for _, offset := range adjacent8 {
					checkRow, checkCol := row+offset[0], col+offset[1]

					if checkRow >= 0 && checkRow < len(grid) &&
						checkCol >= 0 && checkCol < len(grid[checkRow]) &&
						grid[checkRow][checkCol] == '@' {
						adjNeighborsCnt += 1
					}
				}

				if adjNeighborsCnt < 4 {
					recentLoopCnt += 1
					grid[row][col] = 'x'
				}
			}
		}

		count += recentLoopCnt
		if recentLoopCnt == 0 {
			done = true
		}
	}

	return strconv.Itoa(count)
}
