package main

import (
	"aoc/internal/day01"
	"aoc/internal/day02"
	"aoc/internal/day03"
	"aoc/internal/day04"
	"fmt"
)

func main() {
	var days = []func() (string, string){
		day01.Run,
		day02.Run,
		day03.Run,
		day04.Run,
	}

	for i, fun := range days {
		a, b := fun()
		fmt.Printf("[Day %v] %16v :: %16v\n", i+1, a, b)
	}
}
