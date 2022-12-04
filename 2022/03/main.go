package main

import (
	"fmt"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {
	example := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	arr := helper.GetInput("03")
	fmt.Printf("part1 example: %d, should be: %d\n", part1(example), 157)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(example), 70)
	fmt.Printf("part1 input: %d\n", part1(arr))
	fmt.Printf("part2 input: %d\n", part2(arr))

}
func part2(arr []string) int {
	sum := 0
	first := make(map[rune]int)
	second := make(map[rune]int)
	third := make(map[rune]int)
	k := 1
	for _, val := range arr {
		strLen := len(val)
		for i := 0; i < strLen; i++ {
			char := rune(val[i])
			if k == 1 {
				if first[char] == 0 {
					first[char] = 1
				}
			} else if k == 2 {
				if second[char] == 0 {
					second[char] = 1
				}
			} else {
				if third[char] == 0 {
					third[char] = 1
				}
			}
		}
		k++
		if k == 4 {
			for c := range first {
				if second[c] != 0 && third[c] != 0 {
					var lesser int
					if int(c) > 96 {
						lesser = 96
					} else {
						lesser = 38
					}
					sum += int(c) - lesser
				}
			}
			k = 1
			for k2 := range first {
				delete(first, k2)
			}
			for k2 := range second {
				delete(second, k2)
			}
			for k2 := range third {
				delete(third, k2)
			}
		}
	}
	return sum
}
func part1(arr []string) int {
	sum := 0
	for _, val := range arr {
		first := make(map[rune]int)
		second := make(map[rune]int)
		strLen := len(val) / 2
		for idx, char := range val {
			if idx < strLen {
				if first[char] == 0 {
					first[char] = 1
				}
			} else {
				if second[char] == 0 {
					second[char] = 1
				}
			}
		}
		for c := range first {
			if second[c] != 0 {
				var lesser int
				if int(c) > 96 {
					lesser = 96
				} else {
					lesser = 38
				}
				sum += int(c) - lesser
			}
		}
	}
	return sum
}
