package day04

import "os"

func Run() (string, string) {
	data, _ := os.ReadFile("internal/day04/input.txt")
	return Part1(string(data)), Part2(string(data))
}
