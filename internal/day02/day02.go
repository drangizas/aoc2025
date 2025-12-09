package day02

import "os"

func Run() (string, string) {
	data, _ := os.ReadFile("internal/day02/input.txt")
	return Part1(string(data)), Part2(string(data))
}
