package main

import (
	"fmt"
	"strings"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {
	exampleInput := `30373
25512
65332
33549
35390`
	s := strings.Split(exampleInput, "\n")
	l := parseArr(s)
	fmt.Printf("part1 example: %d, should be: %d\n", part1(l), 21)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(l), 8)
	arr := helper.GetInput("08")
	fmt.Printf("part1 input: %d\n", part1(parseArr(arr)))
	fmt.Printf("part2 input: %d\n", part2(parseArr(arr)))
}

func parseArr(arr []string) [][]string {
	l := [][]string{}
	for _, val := range arr {
		str := strings.Split(val, "")
		l = append(l, str)
	}
	return l
}

func part1(strArr [][]string) int {
	invisible := 0
	for i := 1; i < len(strArr); i++ {
		for j := 1; j < len(strArr[i]); j++ {
			isLeftGreater, isRightGreater, isTopGreater, isBottomGreater := false, false, false, false
			z := j
			for z = j - 1; z >= 0; z-- {
				if strArr[i][j] <= strArr[i][z] {
					isLeftGreater = true
					break
				}
			}
			for z := j + 1; z < len(strArr[i]); z++ {
				if strArr[i][j] <= strArr[i][z] {
					isRightGreater = true
					break
				}
			}
			for z := i - 1; z >= 0; z-- {
				if strArr[i][j] <= strArr[z][j] {
					isTopGreater = true
					break
				}
			}
			for z := i + 1; z < len(strArr[i]); z++ {
				if strArr[i][j] <= strArr[z][j] {
					isBottomGreater = true
					break
				}
			}
			if isLeftGreater && isRightGreater && isBottomGreater && isTopGreater {
				invisible++
			}
		}
	}
	return len(strArr)*len(strArr) - invisible
}

func part2(strArr [][]string) int {
	best := 0
	for i := 1; i < len(strArr); i++ {
		for j := 1; j < len(strArr[i]); j++ {
			isLeftGreater, isRightGreater, isTopGreater, isBottomGreater := 0, 0, 0, 0
			for z := j - 1; z >= 0; z-- {
				isLeftGreater++
				if strArr[i][j] <= strArr[i][z] {
					break
				}
			}
			for z := j + 1; z < len(strArr[i]); z++ {
				isRightGreater++
				if strArr[i][j] <= strArr[i][z] {
					break
				}
			}
			for z := i - 1; z >= 0; z-- {
				isTopGreater++
				if strArr[i][j] <= strArr[z][j] {
					break
				}
			}
			for z := i + 1; z < len(strArr[i]); z++ {
				isBottomGreater++
				if strArr[i][j] <= strArr[z][j] {
					break
				}
			}
			gt := isLeftGreater * isBottomGreater * isTopGreater * isRightGreater
			if best <= gt {
				best = gt
			}
		}
	}
	return best
}
