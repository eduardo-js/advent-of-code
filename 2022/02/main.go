package main

import (
	"fmt"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {
	arr := helper.GetInput("02")
	example := []string{"A Y", "B X", "C Z"}
	fmt.Printf("part1 example: %d, should be: %d\n", part1(example), 15)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(example), 12)
	fmt.Printf("part1 input: %d\n", part1(arr))
	fmt.Printf("part2 input: %d\n", part2(arr))
}

const (
	WIN  = 6
	DRAW = 3
)

func part1(arr []string) int {
	var sum int
	for _, val := range arr {
		player := rune(val[2])
		elf := rune(val[0])
		var point int
		if player == 'X' {
			point = 1
		} else if player == 'Y' {
			point = 2
		} else {
			point = 3
		}

		if player == 'Y' && elf == 'A' || player == 'X' && elf == 'C' || player == 'Z' && elf == 'B' {
			sum += WIN
		} else if player == 'X' && elf == 'A' || player == 'Y' && elf == 'B' || player == 'Z' && elf == 'C' {
			sum += DRAW
		}
		sum += point
	}
	return sum
}

func part2(arr []string) int {
	var sum int
	for _, val := range arr {
		player := rune(val[2])
		elf := rune(val[0])

		if player == 'Y' {
			if elf == 'A' {
				sum += 1
			}
			if elf == 'B' {
				sum += 2
			}
			if elf == 'C' {
				sum += 3
			}
			sum += DRAW
		}

		if player == 'Z' {
			if elf == 'A' {
				sum += 2
			}
			if elf == 'B' {
				sum += 3
			}
			if elf == 'C' {
				sum += 1
			}
			sum += WIN
		}

		if player == 'X' {
			if elf == 'A' {
				sum += 3
			}
			if elf == 'B' {
				sum += 1
			}
			if elf == 'C' {
				sum += 2
			}

		}
	}
	return sum
}
