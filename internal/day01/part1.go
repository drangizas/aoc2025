package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	current, rotations := 50, 0

	sc := bufio.NewScanner(strings.NewReader(input))
	for sc.Scan() {
		line := sc.Text()

		dir := line[0]
		n, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(fmt.Sprintf("Failed to parse string from %v", line))
		}

		switch dir {
		case 'L':
			current -= n
		case 'R':
			current += n
		default:
			panic("unknown dir")
		}

		current = ((current % 100) + 100) % 100

		if current == 0 {
			rotations += 1
		}

	}

	return strconv.Itoa(rotations)
}
