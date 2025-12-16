package day04

import (
	"strconv"
	"strings"
)

var adjacent8 = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1}, // top row
	{0, -1}, {0, 1}, // middle row (skip center)
	{1, -1}, {1, 0}, {1, 1}, // bottom row
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	var count int
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
				count += 1
			}
		}
	}

	return strconv.Itoa(count)
}
