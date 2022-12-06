package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {

	example := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	exampleMoves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	exampleCopy := make([][]string, len(example))
	copy(exampleCopy, example)

	fmt.Printf("part1 example: %s, should be: %s\n", part1(example, exampleMoves), "CMZ")
	fmt.Printf("part2 example: %s, should be: %s\n", part2(exampleCopy, exampleMoves), "MCD")
	arr := helper.GetInput("05")
	mt := [][]string{
		{"R", "G", "H", "Q", "S", "B", "T", "N"},
		{"H", "S", "F", "D", "P", "Z", "J"},
		{"Z", "H", "V"},
		{"M", "Z", "J", "F", "G", "H"},
		{"T", "Z", "C", "D", "L", "M", "S", "R"},
		{"M", "T", "W", "V", "H", "Z", "J"},
		{"T", "F", "P", "L", "Z"},
		{"Q", "V", "W", "S"},
		{"W", "H", "L", "M", "T", "D", "N", "C"},
	}
	mtCopy := make([][]string, len(mt))
	copy(mtCopy, mt)

	fmt.Printf("part1 input: %s\n", part1(mt, arr))
	fmt.Printf("part2 input: %s\n", part2(mtCopy, arr))
}
func part1(mt [][]string, arr []string) string {
	for _, moves := range arr {
		str := strings.Split(moves, " ")
		to, _ := strconv.ParseInt(str[5], 10, 0)
		from, _ := strconv.ParseInt(str[3], 10, 0)
		amount, _ := strconv.ParseInt(str[1], 10, 0)
		from--
		to--
		idx := math.Max(float64(len(mt[from])-int(amount)), 0.0)
		copy := mt[from][int(idx):]
		reverse(copy)
		mt[to] = append(mt[to], copy...)
		mt[from] = mt[from][:int(idx)]
	}
	str := ""
	for _, z := range mt {
		str += z[len(z)-1]
	}
	return str
}

func reverse(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func part2(mt [][]string, arr []string) string {
	for _, moves := range arr {
		str := strings.Split(moves, " ")
		to, _ := strconv.ParseInt(str[5], 10, 0)
		from, _ := strconv.ParseInt(str[3], 10, 0)
		amount, _ := strconv.ParseInt(str[1], 10, 0)
		from--
		to--
		idx := math.Max(float64(len(mt[from])-int(amount)), 0.0)
		mt[to] = append(mt[to], mt[from][int(idx):]...)
		mt[from] = mt[from][:int(idx)]
	}
	str := ""
	for _, z := range mt {
		str += z[len(z)-1]
	}
	return str
}
