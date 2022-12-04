package main

import (
	"fmt"
	"strconv"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func part1(arr []string) int {
	arr = append(arr, "")
	var sum int
	var greater int
	for _, value := range arr {

		if value == "" {
			if greater < sum {
				greater = sum
			}
			sum = 0
		} else {
			v, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				panic(err)
			}
			sum += int(v)
		}

	}
	return greater
}

func min(first, second *int) *int {
	if *first < *second {
		return first
	}
	return second
}
func part2(arr []string) int {
	arr = append(arr, "")
	var first, second, third int
	var sum int
	for _, value := range arr {
		if value == "" {
			k := min(min(&first, &second), &third)
			if *k < sum {
				*k = sum
			}
			sum = 0
		} else {
			v, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				panic(err)
			}
			sum += int(v)
		}
	}
	return first + second + third
}

func main() {
	arr := helper.GetInput("01")
	example := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	fmt.Printf("part1 example: %d, should be: %d\n", part1(example), 24000)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(example), 45000)
	fmt.Printf("part1 input: %d\n", part1(arr))
	fmt.Printf("part2 input: %d\n", part2(arr))
}
