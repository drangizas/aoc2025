package day03

import "os"

func Run() (string, string) {
	data, _ := os.ReadFile("internal/day03/input.txt")
	return Part1(string(data)), Part2(string(data))
}
