package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2(input string) string {
	current, rotations := 50, 0
	prevCurrent := current

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

		if current > 0 {
			rotations += current / 100
		} else if current <= 0 {
			toAdd := 0
			if prevCurrent != 0 {
				toAdd = 1
			}

			rotations += (current / -100) + toAdd
		}

		current = ((current % 100) + 100) % 100
		prevCurrent = current

		// fmt.Printf("[%4v] [%4v] %v\n", rotations, current, line)
	}

	return strconv.Itoa(rotations)

}
