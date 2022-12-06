package main

import (
	"fmt"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {

	fmt.Printf("part1 example: %d, should be: %d\n", part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7)
	fmt.Printf("part1 example: %d, should be: %d\n", part1("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5)
	fmt.Printf("part1 example: %d, should be: %d\n", part1("nppdvjthqldpwncqszvftbrmjlhg"), 6)
	fmt.Printf("part1 example: %d, should be: %d\n", part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10)
	fmt.Printf("part1 example: %d, should be: %d\n", part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11)
	fmt.Printf("part2 example: %d, should be: %d\n", part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19)
	fmt.Printf("part2 example: %d, should be: %d\n", part2("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23)
	fmt.Printf("part2 example: %d, should be: %d\n", part2("nppdvjthqldpwncqszvftbrmjlhg"), 23)
	fmt.Printf("part2 example: %d, should be: %d\n", part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29)
	fmt.Printf("part2 example: %d, should be: %d\n", part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26)

	arr := helper.GetInput("06")[0]
	fmt.Printf("part1 input: %d\n", part1(arr))
	fmt.Printf("part2 input: %d\n", part2(arr))
}

func part1(str string) int {
	return ans(str, 4)
}

func part2(str string) int {
	return ans(str, 14)
}

func ans(str string, size int) int {
	for i := range str {
		m := map[byte]byte{}
		m[str[i]] = str[i]
		for k := 1; k < size; k++ {
			_, ok := m[str[i+k]]
			if ok {
				break
			}
			m[str[i+k]] = str[i+k]
		}
		if len(m) == size {
			return i + size
		}
	}
	return -1
}
