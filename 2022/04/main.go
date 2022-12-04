package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {
	example := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
	fmt.Printf("part1 example: %d, should be: %d\n", part1(example), 2)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(example), 4)
	arr := helper.GetInput("04")
	fmt.Printf("part1 input: %d\n", part1(arr))
	fmt.Printf("part2 input: %d\n", part2(arr))
}
func part1(arr []string) int {
	sum := 0
	for _, val := range arr {
		a, b, c, d := getParams(val)

		if int(a) <= int(c) && int(b) >= int(d) || int(c) <= int(a) && int(d) >= int(b) {
			sum++
		}
	}
	return sum
}

func getParams(val string) (a, b, c, d int) {
	strArr := strings.Split(val, ",")

	first := strings.Split(strArr[0], "-")
	second := strings.Split(strArr[1], "-")

	x, _ := strconv.ParseInt(first[0], 10, 16)
	y, _ := strconv.ParseInt(first[1], 10, 16)
	z, _ := strconv.ParseInt(second[0], 10, 16)
	w, _ := strconv.ParseInt(second[1], 10, 16)
	return int(x), int(y), int(z), int(w)
}

func part2(arr []string) int {
	sum := 0
	for _, val := range arr {
		a, b, c, d := getParams(val)

		if int(a) <= int(c) && int(b) >= int(c) || int(c) <= int(a) && int(d) >= int(a) {
			sum++
		}

	}
	return sum
}
